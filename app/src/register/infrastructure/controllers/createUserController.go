package controllers

import (
	"github.com/BryanChanona/api_chombi.git/app/src/register/application/useCase"
	"github.com/BryanChanona/api_chombi.git/app/src/register/domain/entities"
	"github.com/gin-gonic/gin"
)


type CreateUserController struct{
	useCase *useCase.CreateUserUseCase
}


func NewCreateUserController(useCase *useCase.CreateUserUseCase)*CreateUserController{
	return &CreateUserController{useCase: useCase}
}


func (controller *CreateUserController) Execute(ctx *gin.Context){

	var user entities.User
// Bindear el JSON del request al struct User
	if err := ctx.ShouldBindJSON(&user); err != nil {
	ctx.JSON(400, gin.H{
		"error": "JSON inválido",
	})
	return
}
// Validar los datos del usuario
if err := entities.ValidateUser(user); err != nil {
	ctx.JSON(400, gin.H{
		"error": err.Error(),
	})
	return
}
// Ejecutar el caso de uso para crear el usuario
if err := controller.useCase.Execute(user); err != nil {
	ctx.JSON(500, gin.H{
		"error": err.Error(),
	})
	return
}
ctx.JSON(201, gin.H{
	"message": "Usuario creado exitosamente",
})

}