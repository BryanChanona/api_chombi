package useCase

import (

	"github.com/BryanChanona/api_chombi.git/app/src/vehicles/domain/repositories"
)
type DeleteVehicleUseCase struct {
    repo repositories.IVehicleRepository
}
func NewDeleteVehicleUseCase(repo repositories.IVehicleRepository) *DeleteVehicleUseCase {
    return &DeleteVehicleUseCase{repo: repo}
}
func (uc *DeleteVehicleUseCase) Execute(id int, userId int) error {
    return uc.repo.Delete(id, userId)
}