package service

import (
	"github.com/yoshinorihisakawa/sample-api-hoop/domain/model"
	"github.com/yoshinorihisakawa/sample-api-hoop/usecase/presenter"
	"github.com/yoshinorihisakawa/sample-api-hoop/usecase/repository"
)

type userService struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserService interface {
	Create(u *model.User) error
	Get(u []*model.User) ([]*model.User, error)
}

func NewUserService(repo repository.UserRepository, pre presenter.UserPresenter) UserService {
	return &userService{repo, pre}
}

func (userService *userService) Create(u *model.User) error {
	return userService.UserRepository.Store(u)
}

func (userService *userService) Get(u []*model.User) ([]*model.User, error) {
	us, err := userService.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}
	return userService.UserPresenter.ResponseUsers(us), nil
}
