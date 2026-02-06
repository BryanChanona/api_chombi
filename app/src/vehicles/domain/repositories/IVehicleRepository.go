package repositories

import "github.com/BryanChanona/api_chombi.git/app/src/vehicles/domain/entities"

type IVehicleRepository interface {
    Save(vehicle entities.Vehicle) error
    GetAll(userId int) ([]entities.VehicleResponse, error)
    Update(id int, userId int, v entities.Vehicle) error
    Delete(id int, userId int) error
    GetById(id int) (entities.VehicleResponse, error)
}