package presenters

import (
	"github.com/yoshinorihisakawa/sample-api-hoop/domain/model"
)

type userPresenter struct {
}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

type UserPresenter interface {
	ResponseUsers(us []*model.User) []*model.User
}

func (userPresenter *userPresenter) ResponseUsers(us []*model.User) []*model.User {
	for _, u := range us {
		u.LastName = u.LastName + "hoge"
	}
	return us
}
