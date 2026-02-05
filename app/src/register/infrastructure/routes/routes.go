package routes

import (
	"github.com/BryanChanona/api_chombi.git/app/src/register/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine){
	routes := router.Group("/api/v1/register")
	{
		createUserController := dependencies.GetCreateUserController()
		routes.POST("/", createUserController.Execute)
	}

}