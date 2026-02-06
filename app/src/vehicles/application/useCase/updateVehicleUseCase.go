package useCase
import (
	"github.com/BryanChanona/api_chombi.git/app/src/vehicles/domain/entities"
	"github.com/BryanChanona/api_chombi.git/app/src/vehicles/domain/repositories"
)
type UpdateVehicleUseCase struct {
    repo repositories.IVehicleRepository
}
func NewUpdateVehicleUseCase(repo repositories.IVehicleRepository) *UpdateVehicleUseCase {
    return &UpdateVehicleUseCase{repo: repo}
}
func (uc *UpdateVehicleUseCase) Execute(id int, v entities.Vehicle) error {
    return uc.repo.Update(id, v)
}