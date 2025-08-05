package errno

import "github.com/crazyfrankie/gem/gerrors"

var (
	Success = gerrors.NewBizError(20000, "success")

	ErrParams         = gerrors.NewBizError(40000, "params in invalid")
	ErrAuthentication = gerrors.NewBizError(40001, "authentication error")
	ErrUnauthorized   = gerrors.NewBizError(40002, "unauthorized")
	ErrExists         = gerrors.NewBizError(40003, "resource already exists")
	ErrNotFound       = gerrors.NewBizError(40004, "resources not found")
	ErrInternalServer = gerrors.NewBizError(50000, "internal server error")
)
