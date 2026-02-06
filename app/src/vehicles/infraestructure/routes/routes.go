    package routes

    import (
        "github.com/BryanChanona/api_chombi.git/app/src/middlewares"
        "github.com/BryanChanona/api_chombi.git/app/src/vehicles/infraestructure/dependencies"
        "github.com/gin-gonic/gin"
    )

    func VehicleRoutes(router *gin.Engine) {
        routes := router.Group("/api/v1/vehicles")
        
        create := dependencies.GetCreateVehicleController().Create
        getAll := dependencies.GetGetAllVehicleController().View
        delete := dependencies.GetDeleteVehicleController().Delete
        update := dependencies.GetUpdateVehicleController().Update
        
        
        routes.POST("/", middlewares.AuthMiddleware(), create)
        routes.GET("/", middlewares.AuthMiddleware(),getAll)
        routes.DELETE("/:id", middlewares.AuthMiddleware(),delete)
        routes.PUT("/:id", middlewares.AuthMiddleware(), update)
    }