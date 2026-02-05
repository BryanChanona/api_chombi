package main

import (
	"github.com/BryanChanona/api_chombi.git/app/src/helpers"
	registerDependencies "github.com/BryanChanona/api_chombi.git/app/src/register/infrastructure/dependencies"
	registerRoutes "github.com/BryanChanona/api_chombi.git/app/src/register/infrastructure/routes"
	"github.com/gin-gonic/gin"
	
)


func main(){
	//Inicializar dependencias
	registerDependencies.InitDependencies()

	r := gin.Default()
	helpers.InitCORS(r)
	//Inicialozar rutas
	registerRoutes.Routes(r)

	
	r.Run(":8080")


}
