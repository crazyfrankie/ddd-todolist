package appinfra

import (
	"context"
	"os"

	"gorm.io/gorm"

	"github.com/crazyfrankie/ddd-todolist/backend/infra/contract/imagex"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/impl/cache/redis"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/impl/idgen"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/impl/imagex/veimagex"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/impl/mysql"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/impl/storage"
	"github.com/crazyfrankie/ddd-todolist/backend/types/consts"
)

type AppDependencies struct {
	DB           *gorm.DB
	CacheCli     *redis.Client
	IDGenSVC     idgen.IDGenerator
	Storage      storage.Storage
	ImageXClient imagex.ImageX
}

func Init(ctx context.Context) (*AppDependencies, error) {
	deps := &AppDependencies{}
	var err error

	deps.DB, err = mysql.New()
	if err != nil {
		return nil, err
	}

	deps.CacheCli = redis.New()

	deps.IDGenSVC, err = idgen.New(deps.CacheCli)
	if err != nil {
		return nil, err
	}

	deps.ImageXClient, err = initImageX(ctx)
	if err != nil {
		return nil, err
	}

	deps.Storage, err = initStorage(ctx)
	if err != nil {
		return nil, err
	}

	return deps, nil
}

func initImageX(ctx context.Context) (imagex.ImageX, error) {
	uploadComponentType := os.Getenv(consts.FileUploadComponentType)

	if uploadComponentType != consts.FileUploadComponentTypeImagex {
		return storage.NewImagex(ctx)
	}
	return veimagex.New(
		os.Getenv(consts.VeImageXAK),
		os.Getenv(consts.VeImageXSK),
		os.Getenv(consts.VeImageXDomain),
		os.Getenv(consts.VeImageXUploadHost),
		os.Getenv(consts.VeImageXTemplate),
		[]string{os.Getenv(consts.VeImageXServerID)},
	)
}

func initStorage(ctx context.Context) (storage.Storage, error) {
	return storage.New(ctx)
}
