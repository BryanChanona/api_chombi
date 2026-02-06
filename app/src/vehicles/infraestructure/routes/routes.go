package routes

import (
	"github.com/BryanChanona/api_chombi.git/app/src/middlewares"
	"github.com/BryanChanona/api_chombi.git/app/src/vehicles/infraestructure/dependencies"
	"github.com/gin-gonic/gin"
)

func VehicleRoutes(router *gin.Engine) {
    routes := router.Group("/vehicles")
    
    create := dependencies.GetCreateVehicleController().Create
    getAll := dependencies.GetGetAllVehicleController().View
    delete := dependencies.GetDeleteVehicleController().Delete
    update := dependencies.GetUpdateVehicleController().Update
    
    
    routes.POST("/", create,middlewares.AuthMiddleware())
    routes.GET("/", getAll, middlewares.AuthMiddleware())
    routes.DELETE("/:id", delete,middlewares.AuthMiddleware())
    routes.PUT("/:id", update,middlewares.AuthMiddleware())
}