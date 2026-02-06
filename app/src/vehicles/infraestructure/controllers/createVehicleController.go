package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/BryanChanona/api_chombi.git/app/src/vehicles/application/useCase"
	"github.com/BryanChanona/api_chombi.git/app/src/vehicles/domain/entities"
)

type CreateVehicleController struct {
    useCase *useCase.CreateVehicleUseCase
	
} 


func NewCreateVehicleController(useCase *useCase.CreateVehicleUseCase) *CreateVehicleController {
    return &CreateVehicleController{useCase: useCase}
}

func (c *CreateVehicleController) Create(ctx *gin.Context) {
    var vehicle entities.Vehicle
    user_id_str, exist := ctx.Get("user_id")
    if !exist {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
        return
    }
    
    user_id, ok := user_id_str.(int)
    if !ok {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el ID del usuario"})
        return
    }
    
    vehicle.UserId = user_id
    
    if err := ctx.ShouldBindJSON(&vehicle); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    if err := c.useCase.Execute(vehicle); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, gin.H{"message": "Vehículo registrado"})
}