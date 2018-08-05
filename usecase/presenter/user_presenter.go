package presenter

import "github.com/yoshinorihisakawa/sample-api-hoop/domain/model"

type UserPresenter interface {
	ResponseUsers(us []*model.User) []*model.User
}
