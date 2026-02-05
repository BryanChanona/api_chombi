package repositories

import "github.com/BryanChanona/api_chombi.git/app/src/login/domain/entities"


type ILoginRepository interface {
	LogIn(email string) (entities.UserWithPassword, error)
}