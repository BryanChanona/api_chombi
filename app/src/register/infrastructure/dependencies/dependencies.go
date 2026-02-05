package dependencies

import (
	"github.com/BryanChanona/api_chombi.git/app/src/helpers"
	"github.com/BryanChanona/api_chombi.git/app/src/register/application/useCase"
	"github.com/BryanChanona/api_chombi.git/app/src/register/infrastructure"
	"github.com/BryanChanona/api_chombi.git/app/src/register/infrastructure/controllers"
)


var(
	mySQL infrastructure.MySQL
)


func InitDependencies() {
	db, err := helpers.ConectToMySQL()
	if err != nil {
		panic(err)
	}

	mySQL = *infrastructure.NewMySQL(db)
}


func GetCreateUserController() *controllers.CreateUserController{
	createUserUseCase := useCase.NewCreateUserUseCase(&mySQL)
	return controllers.NewCreateUserController(createUserUseCase)
}
