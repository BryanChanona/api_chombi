package routes

import (
	"github.com/BryanChanona/api_chombi.git/app/src/login/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine){
	routes := router.Group("/api/v1/login")
	{
		loginController := dependencies.GetLogInController()
		routes.POST("/", loginController.Execute)
	}

}	