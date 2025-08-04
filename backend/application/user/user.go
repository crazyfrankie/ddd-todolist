package user

import (
	user "github.com/crazyfrankie/ddd-todolist/backend/domain/user/service"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/contract/storage"
)

var UserApplicationSVC = &UserApplicationService{}

type UserApplicationService struct {
	oss       storage.Storage
	DomainSVC user.User
}
