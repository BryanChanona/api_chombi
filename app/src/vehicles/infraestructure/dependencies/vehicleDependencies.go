package dependencies

import (
    "github.com/BryanChanona/api_chombi.git/app/src/helpers"
	
    "github.com/BryanChanona/api_chombi.git/app/src/vehicles/application/useCase"
    "github.com/BryanChanona/api_chombi.git/app/src/vehicles/infraestructure"
    "github.com/BryanChanona/api_chombi.git/app/src/vehicles/infraestructure/controllers"
)

var (
    mySQL infraestructure.MySQL
)

func InitVehicles() {
    db, err := helpers.ConectToMySQL()
	if err != nil {
		panic(err)
	}

	mySQL = *infraestructure.NewMySQL(db)
}

func GetCreateVehicleController() *controllers.CreateVehicleController {
    caseCreate := useCase.NewCreateVehicleUseCase(&mySQL)
    return controllers.NewCreateVehicleController(caseCreate)
}

func GetGetAllVehicleController() *controllers.GetAllVehicleController {
    caseGetAll := useCase.NewGetAllVehiclesUseCase(&mySQL)
    return controllers.NewGetAllVehicleController(caseGetAll)
}

func GetDeleteVehicleController() *controllers.DeleteVehicleController {
    caseDelete := useCase.NewDeleteVehicleUseCase(&mySQL)
    return controllers.NewDeleteVehicleController(caseDelete)
}

func GetUpdateVehicleController() *controllers.UpdateVehicleController {
    caseUpdate := useCase.NewUpdateVehicleUseCase(&mySQL)
    return controllers.NewUpdateVehicleController(caseUpdate)
}

