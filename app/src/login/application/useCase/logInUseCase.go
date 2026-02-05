package useCase

import (
	"errors"

	"github.com/BryanChanona/api_chombi.git/app/src/login/domain/entities"
	"github.com/BryanChanona/api_chombi.git/app/src/login/domain/repositories"
	"github.com/BryanChanona/api_chombi.git/app/src/security/domain"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type LogInUseCase struct {
	db     repositories.ILoginRepository
	hasher domain.PasswordHasher
}

func NewLogInUseCase(
	db repositories.ILoginRepository,
	hasher domain.PasswordHasher,
) *LogInUseCase {
	return &LogInUseCase{db: db, hasher: hasher}
}

func (uc *LogInUseCase) Execute(user entities.User) (entities.UserResponse, error) {
	dbUser, err := uc.db.LogIn(user.Email)
	if err != nil {
		return entities.UserResponse{}, ErrInvalidCredentials
	}

	if !uc.hasher.Compare(dbUser.PasswordHash, user.Password) {
		return entities.UserResponse{}, ErrInvalidCredentials
	}

	return entities.UserResponse{
		Id:       dbUser.Id,
		Name:     dbUser.Name,
		LastName: dbUser.LastName,
		Email:    dbUser.Email,
	}, nil
}
