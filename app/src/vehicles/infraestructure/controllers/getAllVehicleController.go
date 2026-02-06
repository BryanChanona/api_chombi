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
    // 1. Obtener el ID del usuario desde el context
    userIdValue, exists := ctx.Get("user_id")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
        return
    }

    userId, ok := userIdValue.(int)
    if !ok {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno al validar identidad"})
        return
    }

    // 2. Pasar el userId al caso de uso
    vehicles, err := c.useCase.Execute(userId)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 3. Responder con los datos filtrados
    ctx.JSON(http.StatusOK, gin.H{"vehicles": vehicles})
}