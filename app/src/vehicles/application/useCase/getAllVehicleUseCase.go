package useCase

import (
	"github.com/BryanChanona/api_chombi.git/app/src/vehicles/domain/entities"
	"github.com/BryanChanona/api_chombi.git/app/src/vehicles/domain/repositories"
)

type GetAllVehiclesUseCase struct {
    repo repositories.IVehicleRepository
}
func NewGetAllVehiclesUseCase(repo repositories.IVehicleRepository) *GetAllVehiclesUseCase {
    return &GetAllVehiclesUseCase{repo: repo}
}
func (uc *GetAllVehiclesUseCase) Execute() ([]entities.VehicleResponse, error) {
    return uc.repo.GetAll()
}