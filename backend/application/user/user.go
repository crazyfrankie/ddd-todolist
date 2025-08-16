package user

import (
	"context"
	"errors"
	"net/mail"

	model "github.com/crazyfrankie/ddd-todolist/backend/api/model/user"
	"github.com/crazyfrankie/ddd-todolist/backend/application/base/ctxutil"
	"github.com/crazyfrankie/ddd-todolist/backend/domain/user/entity"
	user "github.com/crazyfrankie/ddd-todolist/backend/domain/user/service"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/contract/storage"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/contract/token"
	"github.com/crazyfrankie/ddd-todolist/backend/pkg/ptr"
)

type UserApplicationService struct {
	oss       storage.Storage
	jwtGen    token.JWT
	DomainSVC user.User
}

// Add a simple email verification function
func isValidEmail(email string) bool {
	// If the email string is not in the correct format, it will return an error.
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (u *UserApplicationService) UserRegister(ctx context.Context, ua string, req *model.EmailRegisterRequest) (
	resp *model.User, tokens []string, err error,
) {
	// Verify that the email format is legitimate
	if !isValidEmail(req.Email) {
		return nil, nil, errors.New("invalid email")
	}

	userInfo, err := u.DomainSVC.Create(ctx, &user.CreateUserRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, nil, err
	}

	userInfo, err = u.DomainSVC.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, nil, err
	}

	tokens, err = u.jwtGen.GenerateToken(userInfo.UserID, ua)
	if err != nil {
		return nil, nil, err
	}

	return userDo2PassportTo(userInfo), tokens, nil
}

func (u *UserApplicationService) UserLogin(ctx context.Context, ua string, req *model.EmailLoginRequest) (
	resp *model.User, tokens []string, err error,
) {
	userInfo, err := u.DomainSVC.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, nil, err
	}

	tokens, err = u.jwtGen.GenerateToken(userInfo.UserID, ua)
	if err != nil {
		return nil, nil, err
	}

	return userDo2PassportTo(userInfo), tokens, nil
}

func (u *UserApplicationService) UserLogout(ctx context.Context, ua string) (err error) {
	uid := ctxutil.MustGetUIDFromCtx(ctx)

	err = u.jwtGen.CleanToken(ctx, uid, ua)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserApplicationService) UpdateUserAvatar(ctx context.Context, mimeType string, req *model.UpdateAvatarRequest) (url string, err error) {
	// Get file suffix by MIME type
	var ext string
	switch mimeType {
	case "image/jpeg", "image/jpg":
		ext = "jpg"
	case "image/png":
		ext = "png"
	case "image/gif":
		ext = "gif"
	case "image/webp":
		ext = "webp"
	default:
		return "", errors.New("unsupported image type")
	}

	uid := ctxutil.MustGetUIDFromCtx(ctx)

	url, err = u.DomainSVC.UpdateAvatar(ctx, uid, ext, req.Avatar)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (u *UserApplicationService) UpdateUserProfile(ctx context.Context, req *model.UpdateProfileRequest) (err error) {
	uid := ctxutil.MustGetUIDFromCtx(ctx)

	err = u.DomainSVC.UpdateProfile(ctx, &user.UpdateProfileRequest{
		UserID:     uid,
		Name:       req.Name,
		UniqueName: req.UserUniqueName,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *UserApplicationService) ResetUserPassword(ctx context.Context, req *model.ResetUserPassword) (err error) {
	err = u.DomainSVC.ResetPassword(ctx, req.Email, req.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserApplicationService) GetUserInfo(ctx context.Context) (
	resp *model.User, err error,
) {
	userID := ctxutil.MustGetUIDFromCtx(ctx)

	userInfo, err := u.DomainSVC.GetUserInfo(ctx, userID)
	if err != nil {
		return nil, err
	}

	return userDo2PassportTo(userInfo), nil
}

func userDo2PassportTo(userDo *entity.User) *model.User {
	return &model.User{
		UserID:         userDo.UserID,
		Name:           userDo.Name,
		ScreenName:     ptr.Of(userDo.Name),
		UserUniqueName: userDo.UniqueName,
		Email:          userDo.Email,
		AvatarURL:      userDo.IconURL,
		UserCreateTime: userDo.CreatedAt / 1000,
	}
}
