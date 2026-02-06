package repositories

import "github.com/BryanChanona/api_chombi.git/app/src/vehicles/domain/entities"

type IVehicleRepository interface {
    Save(vehicle entities.Vehicle) error
    GetAll() ([]entities.VehicleResponse, error)
    Update(id int, vehicle entities.Vehicle) error
    Delete(id int) error
    GetById(id int) (entities.VehicleResponse, error)
}