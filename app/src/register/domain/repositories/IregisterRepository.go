package repositories

import "github.com/BryanChanona/api_chombi.git/app/src/register/domain/entities"

type IRegisterRepository interface {
	Save(user entities.User) error
}