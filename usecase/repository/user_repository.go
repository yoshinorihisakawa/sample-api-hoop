package repository

import "github.com/yoshinorihisakawa/sample-api-hoop/domain/model"

type UserRepository interface {
	Store(user *model.User) error
	FindAll(users []*model.User) ([]*model.User, error)
}
