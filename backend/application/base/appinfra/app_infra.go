package appinfra

import (
	"context"
	"gorm.io/gorm"

	"github.com/crazyfrankie/ddd-todolist/backend/conf"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/impl/cache/redis"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/impl/idgen"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/impl/mysql"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/impl/storage"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/impl/token"
)

type AppDependencies struct {
	DB       *gorm.DB
	CacheCli *redis.Client
	JWTGen   token.JWT
	IDGenSVC idgen.IDGenerator
	Storage  storage.Storage
}

func Init(ctx context.Context) (*AppDependencies, error) {
	deps := &AppDependencies{}
	var err error

	deps.DB, err = mysql.New()
	if err != nil {
		return nil, err
	}

	deps.CacheCli = redis.New()

	deps.JWTGen = token.New(deps.CacheCli, conf.GetConf().JWT.SignAlgo, conf.GetConf().JWT.SecretKey)

	deps.IDGenSVC, err = idgen.New(deps.CacheCli)
	if err != nil {
		return nil, err
	}

	deps.Storage, err = storage.New(ctx)
	if err != nil {
		return nil, err
	}

	return deps, nil
}
