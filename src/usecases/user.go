package usecases

import (
	"to-do-list-api/src/entities"
	"to-do-list-api/src/controllers"
)

type UserUseCase struct {
	repo controllers.UserRespository
}

func NewUserUseCase(repo controllers.UserRespository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) CreateUser(user entities.User) error {
	return uc.repo.CreateUser(user)
}