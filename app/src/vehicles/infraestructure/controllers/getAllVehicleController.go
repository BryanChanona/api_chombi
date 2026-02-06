package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/BryanChanona/api_chombi.git/app/src/vehicles/application/useCase"
)

type GetAllVehicleController struct {
    useCase *useCase.GetAllVehiclesUseCase
}

func NewGetAllVehicleController(useCase *useCase.GetAllVehiclesUseCase) *GetAllVehicleController {
    return &GetAllVehicleController{useCase: useCase}
}

func (c *GetAllVehicleController) View(ctx *gin.Context) {
    vehicles, err := c.useCase.Execute()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"vehicles": vehicles})
}