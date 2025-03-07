package entities

type User struct {
	Id int
	Email string
	Hash string
}

// type UserRespository interface {
// 	CreateUser(user User) (User, error)
// }