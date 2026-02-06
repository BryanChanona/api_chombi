    package infraestructure

    import (
        "database/sql"
        "fmt"
        "github.com/BryanChanona/api_chombi.git/app/src/vehicles/domain/entities"
    )

    type MySQL struct {
        DB *sql.DB
    }

    func NewMySQL(db *sql.DB) *MySQL {
        return &MySQL{DB: db}
    }

    func (mysql *MySQL) Save(v entities.Vehicle) error {
        query := "INSERT INTO vehicles (licensePlate, driver, unitNumber, shift, isWorking, userId) VALUES (?, ?, ?, ?, ?, ?)"
        _, err := mysql.DB.Exec(query, v.LicencePlate, v.Driver, v.UnitNumber, v.Shift, v.IsWorking, v.UserId)
        if err != nil {
            return fmt.Errorf("[MySQL] Error al guardar vehículo: %w", err)
        }
        return nil
    }

    func (mysql *MySQL) GetAll(userId int) ([]entities.VehicleResponse, error) {

    query := "SELECT id, licensePlate, driver, unitNumber, shift, isWorking FROM vehicles WHERE userId = ?"
    rows, err := mysql.DB.Query(query, userId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var vehicles []entities.VehicleResponse
    for rows.Next() {
        var v entities.VehicleResponse
        if err := rows.Scan(&v.Id, &v.LicencePlate, &v.Driver, &v.UnitNumber, &v.Shift, &v.IsWorking); err != nil {
            return nil, err
        }
        vehicles = append(vehicles, v)
    }
    
    if vehicles == nil {
        vehicles = []entities.VehicleResponse{}
    }
    
    return vehicles, nil
}

   func (mysql *MySQL) Update(id int, userId int, v entities.Vehicle) error {
    // Agregamos AND userId = ? para asegurar la propiedad
    query := "UPDATE vehicles SET licensePlate=?, driver=?, unitNumber=?, shift=?, isWorking=? WHERE id=? AND userId=?"
    result, err := mysql.DB.Exec(query, v.LicencePlate, v.Driver, v.UnitNumber, v.Shift, v.IsWorking, id, userId)
    if err != nil {
        return fmt.Errorf("[MySQL] Error al actualizar: %w", err)
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil || rowsAffected == 0 {
        // Si rowsAffected es 0, el vehículo no existe o el userId no coincide
        return fmt.Errorf("vehículo no encontrado o no tienes permiso para editarlo")
    }
    return nil
}

    func (mysql *MySQL) Delete(id int, userId int) error {
    query := "DELETE FROM vehicles WHERE id = ? AND userId = ?"
    result, err := mysql.DB.Exec(query, id, userId)
    if err != nil {
        return err
    }
    
    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return fmt.Errorf("no se encontró el vehículo o no tienes permiso para eliminarlo")
    }
    return nil
}

    func (mysql *MySQL) GetById(id int) (entities.VehicleResponse, error) {
        var v entities.VehicleResponse
        query := "SELECT id, licencePlate, driver, unitNumber, shift, isWorking FROM vehicles WHERE id=?"
        row := mysql.DB.QueryRow(query, id)
        
        err := row.Scan(&v.Id, &v.LicencePlate, &v.Driver, &v.UnitNumber, &v.Shift, &v.IsWorking)
        if err != nil {
            if err == sql.ErrNoRows {
                return v, fmt.Errorf("vehículo no encontrado")
            }
            return v, err
        }
        return v, nil
    }