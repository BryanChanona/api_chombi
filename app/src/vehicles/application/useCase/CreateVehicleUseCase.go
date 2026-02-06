package useCase

import (
	"github.com/BryanChanona/api_chombi.git/app/src/vehicles/domain/entities"
	"github.com/BryanChanona/api_chombi.git/app/src/vehicles/domain/repositories"
)

type CreateVehicleUseCase struct {
    repo repositories.IVehicleRepository
}
func NewCreateVehicleUseCase(repo repositories.IVehicleRepository) *CreateVehicleUseCase {
    return &CreateVehicleUseCase{repo: repo}
}
func (uc *CreateVehicleUseCase) Execute(v entities.Vehicle) error {
    return uc.repo.Save(v)
}