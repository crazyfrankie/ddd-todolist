package service

import (
	"context"

	"github.com/crazyfrankie/ddd-todolist/backend/domain/user/entity"
	"github.com/crazyfrankie/ddd-todolist/backend/domain/user/repository"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/contract/idgen"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/contract/storage"
)

type Components struct {
	IconOSS  storage.Storage
	IDGen    idgen.IDGenerator
	UserRepo repository.UserRepository
}

type userImpl struct {
	*Components
}

func NewUserDomain(ctx context.Context, c *Components) User {
	return &userImpl{
		Components: c,
	}
}

func (u *userImpl) Create(ctx context.Context, req *CreateUserRequest) (user *entity.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userImpl) Login(ctx context.Context, email, password string) (user *entity.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userImpl) Logout(ctx context.Context, userID int64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userImpl) ResetPassword(ctx context.Context, email, password string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userImpl) GetUserInfo(ctx context.Context, userID int64) (user *entity.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userImpl) UpdateAvatar(ctx context.Context, userID int64, ext string, imagePayload []byte) (url string, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userImpl) UpdateProfile(ctx context.Context, req *UpdateProfileRequest) (err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userImpl) GetUserProfiles(ctx context.Context, userID int64) (user *entity.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userImpl) MGetUserProfiles(ctx context.Context, userIDs []int64) (users []*entity.User, err error) {
	//TODO implement me
	panic("implement me")
}
