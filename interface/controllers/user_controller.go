package controllers

import (
	"github.com/yoshinorihisakawa/sample-api-hoop/domain/model"
	"github.com/yoshinorihisakawa/sample-api-hoop/usecase/service"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	CreateUser(user *model.User) error
	GetUsers() ([]*model.User, error)
}

func NewUserController(us service.UserService) UserController {
	return &userController{us}
}

func (userController *userController) CreateUser(user *model.User) error {
	// 内側の処理のための技術的な関心事を記載
	return userController.userService.Create(user)
}

func (userController *userController) GetUsers() ([]*model.User, error) {
	u := []*model.User{}
	us, err := userController.userService.Get(u)
	if err != nil {
		return nil, err
	}
	return us, nil
}
