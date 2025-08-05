package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/crazyfrankie/ddd-todolist/backend/infra/contract/storage"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/impl/storage/minio"
	"github.com/crazyfrankie/ddd-todolist/backend/types/consts"
)

type Storage = storage.Storage

func New(ctx context.Context) (Storage, error) {
	storageType := os.Getenv(consts.StorageType)
	switch storageType {
	case "minio":
		return minio.New(
			ctx,
			os.Getenv(consts.MinIOEndpoint),
			os.Getenv(consts.MinIOAK),
			os.Getenv(consts.MinIOSK),
			os.Getenv(consts.StorageBucket),
			false,
		)
	}

	return nil, fmt.Errorf("unknown storage type: %s", storageType)
}
