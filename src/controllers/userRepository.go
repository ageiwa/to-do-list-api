package controllers

import (
	"to-do-list-api/src/entities"
)

type UserRespository interface {
	CreateUser(user entities.User) error
}