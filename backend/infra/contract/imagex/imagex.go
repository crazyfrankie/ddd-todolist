package imagex

import (
	"context"
	"time"
)

type ImageX interface {
	GetUploadAuth(ctx context.Context, opt ...UploadAuthOpt) (*SecurityToken, error)
	GetUploadAuthWithExpire(ctx context.Context, expire time.Duration, opt ...UploadAuthOpt) (*SecurityToken, error)
	GetResourceURL(ctx context.Context, uri string, opts ...GetResourceOpt) (*ResourceURL, error)
	Upload(ctx context.Context, data []byte, opts ...UploadAuthOpt) (*UploadResult, error)
	GetServerID() string
	GetUploadHost(ctx context.Context) string
}

type SecurityToken struct {
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	SessionToken    string `json:"session_token"`
	ExpiredTime     string `json:"expired_time"`
	CurrentTime     string `json:"current_time"`
	HostScheme      string ` json:"host_scheme"`
}

type ResourceURL struct {
	// REQUIRED; The resulting graph accesses the thin address, missing the bucket part compared to the default address.
	CompactURL string `json:"CompactURL"`
	// REQUIRED; Result graph access default address.
	URL string `json:"URL"`
}

type UploadResult struct {
	Result    *Result   `json:"Results"`
	RequestId string    `json:"RequestId"`
	FileInfo  *FileInfo `json:"PluginResult"`
}

type Result struct {
	Uri       string `json:"Uri"`
	UriStatus int    `json:"UriStatus"` // 2000 means the upload was successful.
}

type FileInfo struct {
	Name        string `json:"FileName"`
	Uri         string `json:"ImageUri"`
	ImageWidth  int    `json:"ImageWidth"`
	ImageHeight int    `json:"ImageHeight"`
	Md5         string `json:"ImageMd5"`
	ImageFormat string `json:"ImageFormat"`
	ImageSize   int    `json:"ImageSize"`
	FrameCnt    int    `json:"FrameCnt"`
	Duration    int    `json:"Duration"`
}
