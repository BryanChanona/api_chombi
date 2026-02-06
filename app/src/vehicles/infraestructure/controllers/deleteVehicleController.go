package controllers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "github.com/BryanChanona/api_chombi.git/app/src/vehicles/application/useCase"
)

type DeleteVehicleController struct {
    useCase *useCase.DeleteVehicleUseCase
}

func NewDeleteVehicleController(useCase *useCase.DeleteVehicleUseCase) *DeleteVehicleController {
    return &DeleteVehicleController{useCase: useCase}
}

func (c *DeleteVehicleController) Delete(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }
    
    if err := c.useCase.Execute(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Vehículo eliminado"})
}