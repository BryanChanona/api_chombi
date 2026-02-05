package dependencies

import (
	"github.com/BryanChanona/api_chombi.git/app/src/helpers"
	"github.com/BryanChanona/api_chombi.git/app/src/login/application/useCase"
	"github.com/BryanChanona/api_chombi.git/app/src/login/infrastructure"
	infrastructureSecurity "github.com/BryanChanona/api_chombi.git/app/src/security/infrastructure"

	"github.com/BryanChanona/api_chombi.git/app/src/login/infrastructure/controllers"
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

func GetLogInController() *controllers.LogInController{
	hasher := infrastructureSecurity.NewBCryptHasher()
	useCase := useCase.NewLogInUseCase(&mySQL, hasher)
	return controllers.NewLogInController(useCase)
}