package controllers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "github.com/BryanChanona/api_chombi.git/app/src/vehicles/application/useCase"
    "github.com/BryanChanona/api_chombi.git/app/src/vehicles/domain/entities"
)

type UpdateVehicleController struct {
    useCase *useCase.UpdateVehicleUseCase
}

func NewUpdateVehicleController(useCase *useCase.UpdateVehicleUseCase) *UpdateVehicleController {
    return &UpdateVehicleController{useCase: useCase}
}

func (c *UpdateVehicleController) Update(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    var v entities.Vehicle
    if err := ctx.ShouldBindJSON(&v); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.useCase.Execute(id, v); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Vehículo actualizado"})
}