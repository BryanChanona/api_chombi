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
    // 1. Obtener ID del vehículo de la URL
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de vehículo inválido"})
        return
    }

    // 2. Obtener el ID del usuario desde el context (inyectado por el Middleware)
    userIdValue, exists := ctx.Get("user_id")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
        return
    }

    // Asegurarse de que sea el tipo correcto (int)
    userId, ok := userIdValue.(int)
    if !ok {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno al validar usuario"})
        return
    }

    // 3. Ejecutar la eliminación pasando ambos IDs
    if err := c.useCase.Execute(id, userId); err != nil {
        // Podrías diferenciar entre error de DB o error de "No autorizado"
        ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Vehículo eliminado correctamente"})
}