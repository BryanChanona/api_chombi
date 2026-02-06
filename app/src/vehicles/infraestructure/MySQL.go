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
    query := "INSERT INTO vehicles (licensePlate, driver, unitNumber, shift, isWorking) VALUES (?, ?, ?, ?, ?)"
    _, err := mysql.DB.Exec(query, v.LicencePlate, v.Driver, v.UnitNumber, v.Shift, v.IsWorking)
    if err != nil {
        return fmt.Errorf("[MySQL] Error al guardar vehículo: %w", err)
    }
    return nil
}

func (mysql *MySQL) GetAll() ([]entities.VehicleResponse, error) {
    rows, err := mysql.DB.Query("SELECT id, licensePlate, driver, unitNumber, shift, isWorking FROM vehicles")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var vehicles []entities.VehicleResponse
    for rows.Next() {
        var v entities.VehicleResponse
        // Asegúrate de que el orden coincida con el SELECT
        if err := rows.Scan(&v.Id, &v.LicencePlate, &v.Driver, &v.UnitNumber, &v.Shift, &v.IsWorking); err != nil {
            return nil, err
        }
        vehicles = append(vehicles, v)
    }
    return vehicles, nil
}

func (mysql *MySQL) Update(id int, v entities.Vehicle) error {
    query := "UPDATE vehicles SET licensePlate=?, driver=?, unitNumber=?, shift=?, isWorking=? WHERE id=?"
    result, err := mysql.DB.Exec(query, v.LicencePlate, v.Driver, v.UnitNumber, v.Shift, v.IsWorking, id)
    if err != nil {
        return fmt.Errorf("[MySQL] Error al actualizar: %w", err)
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil || rowsAffected == 0 {
        return fmt.Errorf("no se encontró el vehículo para actualizar")
    }
    return nil
}

func (mysql *MySQL) Delete(id int) error {
    query := "DELETE FROM vehicles WHERE id=?"
    result, err := mysql.DB.Exec(query, id)
    if err != nil {
        return err
    }
    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return fmt.Errorf("no se encontró el vehículo con ID: %d", id)
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