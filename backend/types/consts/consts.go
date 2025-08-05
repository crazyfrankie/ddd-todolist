package consts

import "time"

const (
	SessionDataKeyInCtx = "session_data_key_in_ctx"

	StorageType   = "STORAGE_TYPE"
	MinIOAK       = "MINIO_AK"
	MinIOSK       = "MINIO_SK"
	MinIOEndpoint = "MINIO_ENDPOINT"
	StorageBucket = "STORAGE_BUCKET"
)

const (
	SessionMaxAgeSecond    = 30 * 24 * 60 * 60
	DefaultSessionDuration = SessionMaxAgeSecond * time.Second
)
