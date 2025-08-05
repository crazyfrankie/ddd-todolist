package user

import (
	"context"

	"gorm.io/gorm"

	"github.com/crazyfrankie/ddd-todolist/backend/domain/user/repository"
	"github.com/crazyfrankie/ddd-todolist/backend/domain/user/service"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/contract/idgen"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/contract/storage"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/contract/token"
)

func InitService(ctx context.Context, db *gorm.DB, oss storage.Storage, idgen idgen.IDGenerator, jwtGen token.JWT) *UserApplicationService {
	user := &UserApplicationService{}

	user.DomainSVC = service.NewUserDomain(ctx, &service.Components{
		IconOSS:  oss,
		IDGen:    idgen,
		UserRepo: repository.NewUserRepo(db),
	})
	user.oss = oss
	user.jwtGen = jwtGen

	return user
}
