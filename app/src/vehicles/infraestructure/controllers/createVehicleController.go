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
    var v entities.Vehicle
    if err := ctx.ShouldBindJSON(&v); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    if err := c.useCase.Execute(v); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, gin.H{"message": "Vehículo registrado"})
}