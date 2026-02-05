package useCase

import (
	"github.com/BryanChanona/api_chombi.git/app/src/register/domain/entities"
	"github.com/BryanChanona/api_chombi.git/app/src/register/domain/repositories"
	hash "github.com/BryanChanona/api_chombi.git/app/src/security/infrastructure"
)

type CreateUserUseCase struct {
	db repositories.IRegisterRepository
}

func NewCreateUserUseCase(db repositories.IRegisterRepository)*CreateUserUseCase{
	return &CreateUserUseCase{db: db}
}

func (useCase *CreateUserUseCase) Execute(user entities.User) error {

	hashedPassword, err := hash.NewBCryptHasher().Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return useCase.db.Save(user)
}
