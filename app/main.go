package main

import (
	"github.com/BryanChanona/api_chombi.git/app/src/helpers"
	logInDependencies "github.com/BryanChanona/api_chombi.git/app/src/login/infrastructure/dependencies"
	logInRoutes "github.com/BryanChanona/api_chombi.git/app/src/login/infrastructure/routes"
	registerDependencies "github.com/BryanChanona/api_chombi.git/app/src/register/infrastructure/dependencies"
	registerRoutes "github.com/BryanChanona/api_chombi.git/app/src/register/infrastructure/routes"
	vehicleDependencies "github.com/BryanChanona/api_chombi.git/app/src/vehicles/infraestructure/dependencies"
	vehicleRoutes "github.com/BryanChanona/api_chombi.git/app/src/vehicles/infraestructure/routes"

	"github.com/gin-gonic/gin"
)


func main(){
	//Inicializar dependencias
	registerDependencies.InitDependencies()
	logInDependencies.InitDependencies()
	vehicleDependencies.InitVehicles()



	r := gin.Default()
	helpers.InitCORS(r)
	//Inicialozar rutas
	registerRoutes.Routes(r)
	logInRoutes.Routes(r)
	vehicleRoutes.VehicleRoutes(r)
	
	r.Run(":8080")


}
