package consts

import "time"

const (
	VeImageXAK         = "VE_IMAGEX_AK"
	VeImageXSK         = "VE_IMAGEX_SK"
	VeImageXServerID   = "VE_IMAGEX_SERVER_ID"
	VeImageXDomain     = "VE_IMAGEX_DOMAIN"
	VeImageXTemplate   = "VE_IMAGEX_TEMPLATE"
	VeImageXUploadHost = "VE_IMAGEX_UPLOAD_HOST"

	FileUploadComponentType       = "FILE_UPLOAD_COMPONENT_TYPE"
	FileUploadComponentTypeImagex = "imagex"

	SessionDataKeyInCtx = "session_data_key_in_ctx"

	StorageType        = "STORAGE_TYPE"
	MinIOAK            = "MINIO_AK"
	MinIOSK            = "MINIO_SK"
	MinIOEndpoint      = "MINIO_ENDPOINT"
	MinIOProxyEndpoint = "MINIO_PROXY_ENDPOINT"
	MinIOAPIHost       = "MINIO_API_HOST"
	StorageBucket      = "STORAGE_BUCKET"

	HostKeyInCtx          = "HOST_KEY_IN_CTX"
	RequestSchemeKeyInCtx = "REQUEST_SCHEME_IN_CTX"
)

const (
	ApplyUploadActionURI = "/api/common/upload/apply_upload_action"
	UploadURI            = "/api/common/upload"
)

const (
	SessionMaxAgeSecond    = 30 * 24 * 60 * 60
	DefaultSessionDuration = SessionMaxAgeSecond * time.Second
)
