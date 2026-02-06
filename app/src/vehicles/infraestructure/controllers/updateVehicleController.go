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
    // 1. Obtener ID del vehículo del parámetro
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    // 2. Obtener el ID del usuario autenticado (desde el Middleware)
    userIdValue, exists := ctx.Get("user_id")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
        return
    }
    userId := userIdValue.(int)

    // 3. Bind del JSON con los nuevos datos
    var v entities.Vehicle
    if err := ctx.ShouldBindJSON(&v); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 4. Ejecutar actualización pasando id del vehículo Y id del usuario
    if err := c.useCase.Execute(id, userId, v); err != nil {
        ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Vehículo actualizado correctamente"})
}