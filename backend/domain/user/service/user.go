package service

import (
	"context"

	"github.com/crazyfrankie/ddd-todolist/backend/domain/user/entity"
)

type CreateUserRequest struct {
	Email       string
	Password    string
	Name        string
	UniqueName  string
	Description string
}

type UpdateProfileRequest struct {
	UserID     int64
	Name       *string
	UniqueName *string
}

type ValidateProfileUpdateRequest struct {
	UniqueName *string
	Email      *string
}

type ValidateProfileUpdateResult int

const (
	ValidateSuccess             ValidateProfileUpdateResult = 0
	UniqueNameExist             ValidateProfileUpdateResult = 2
	UniqueNameTooShortOrTooLong ValidateProfileUpdateResult = 3
	EmailExist                  ValidateProfileUpdateResult = 5
)

type ValidateProfileUpdateResponse struct {
	Code ValidateProfileUpdateResult
	Msg  string
}

type User interface {
	// Create creates or registers a new user.
	Create(ctx context.Context, req *CreateUserRequest) (user *entity.User, err error)
	Login(ctx context.Context, email, password string) (user *entity.User, err error)
	ResetPassword(ctx context.Context, email, password string) (err error)
	GetUserInfo(ctx context.Context, userID int64) (user *entity.User, err error)
	UpdateAvatar(ctx context.Context, userID int64, ext string, imagePayload []byte) (url string, err error)
	UpdateProfile(ctx context.Context, req *UpdateProfileRequest) (err error)
	GetUserProfiles(ctx context.Context, userID int64) (user *entity.User, err error)
	MGetUserProfiles(ctx context.Context, userIDs []int64) (users []*entity.User, err error)
}
