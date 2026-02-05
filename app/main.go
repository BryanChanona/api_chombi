package main

import (
	"github.com/BryanChanona/api_chombi.git/app/src/helpers"
	registerDependencies "github.com/BryanChanona/api_chombi.git/app/src/register/infrastructure/dependencies"
	logInDependencies "github.com/BryanChanona/api_chombi.git/app/src/login/infrastructure/dependencies"
	registerRoutes "github.com/BryanChanona/api_chombi.git/app/src/register/infrastructure/routes"
	logInRoutes "github.com/BryanChanona/api_chombi.git/app/src/login/infrastructure/routes"
	"github.com/gin-gonic/gin"
	
	
)


func main(){
	//Inicializar dependencias
	registerDependencies.InitDependencies()
	logInDependencies.InitDependencies()


	r := gin.Default()
	helpers.InitCORS(r)
	//Inicialozar rutas
	registerRoutes.Routes(r)
	logInRoutes.Routes(r)

	
	r.Run(":8080")


}
