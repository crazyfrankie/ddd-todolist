package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/crazyfrankie/frx/errorx"
	"golang.org/x/crypto/bcrypt"

	uploadEntity "github.com/crazyfrankie/ddd-todolist/backend/domain/upload/entity"
	"github.com/crazyfrankie/ddd-todolist/backend/domain/user/entity"
	userEntity "github.com/crazyfrankie/ddd-todolist/backend/domain/user/entity"
	"github.com/crazyfrankie/ddd-todolist/backend/domain/user/internal/dal/model"
	"github.com/crazyfrankie/ddd-todolist/backend/domain/user/repository"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/contract/idgen"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/contract/storage"
	"github.com/crazyfrankie/ddd-todolist/backend/pkg/logs"
	"github.com/crazyfrankie/ddd-todolist/backend/pkg/ptr"
	"github.com/crazyfrankie/ddd-todolist/backend/types/errno"
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
	exist, err := u.UserRepo.CheckEmailExist(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if exist {
		return nil, errorx.New(errno.ErrEmailExistCode)
	}

	if req.UniqueName != "" {
		exist, err = u.UserRepo.CheckUniqueNameExist(ctx, req.UniqueName)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, errorx.New(errno.ErrUniqueNameExistCode)
		}
	}

	hashedPasswd, err := hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	name := req.Name
	if name == "" {
		name = strings.Split(req.Email, "@")[0]
	}

	userID, err := u.IDGen.GenID(ctx)
	if err != nil {
		return nil, fmt.Errorf("generate id error: %v", err)
	}

	now := time.Now().UnixMilli()

	newUser := &model.User{
		ID:         userID,
		Name:       name,
		UniqueName: u.getUniqueNameFormEmail(ctx, req.Email),
		Email:      req.Email,
		Password:   hashedPasswd,
		IconURI:    uploadEntity.UserIconURI,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	err = u.UserRepo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, fmt.Errorf("insert user failed: %w", err)
	}

	iconURL, err := u.IconOSS.GetObjectUrl(ctx, newUser.IconURI)
	if err != nil {
		return nil, fmt.Errorf("get icon url failed: %w", err)
	}

	return userPo2Do(newUser, iconURL), nil
}

func (u *userImpl) getUniqueNameFormEmail(ctx context.Context, email string) string {
	arr := strings.Split(email, "@")
	if len(arr) != 2 {
		return email
	}

	username := arr[0]

	exist, err := u.UserRepo.CheckUniqueNameExist(ctx, username)
	if err != nil {
		logs.CtxWarnf(ctx, "check unique name exist failed: %v", err)
		return email
	}

	if exist {
		logs.CtxWarnf(ctx, "unique name %s already exist", username)

		return email
	}

	return username
}

func (u *userImpl) Login(ctx context.Context, email, password string) (user *entity.User, err error) {
	userModel, exist, err := u.UserRepo.GetUsersByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errorx.New(errno.ErrUserNotFoundCode)
	}

	err = verifyPassword(password, userModel.Password)
	if err != nil {
		return nil, err
	}

	resURL, err := u.IconOSS.GetObjectUrl(ctx, userModel.IconURI)
	if err != nil {
		return nil, err
	}

	return userPo2Do(userModel, resURL), nil
}

func (u *userImpl) ResetPassword(ctx context.Context, email, password string) (err error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	err = u.UserRepo.UpdatePassword(ctx, email, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func (u *userImpl) GetUserInfo(ctx context.Context, userID int64) (user *entity.User, err error) {
	if userID <= 0 {
		return nil, fmt.Errorf("invalid user id")
	}

	userModel, err := u.UserRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	resURL, err := u.IconOSS.GetObjectUrl(ctx, userModel.IconURI)
	if err != nil {
		return nil, err
	}

	return userPo2Do(userModel, resURL), nil
}

func (u *userImpl) UpdateAvatar(ctx context.Context, userID int64, ext string, imagePayload []byte) (url string, err error) {
	avatarKey := "user_avatar/" + strconv.FormatInt(userID, 10) + "." + ext
	err = u.IconOSS.PutObject(ctx, avatarKey, imagePayload)
	if err != nil {
		return "", err
	}

	err = u.UserRepo.UpdateAvatar(ctx, userID, avatarKey)
	if err != nil {
		return "", err
	}

	url, err = u.IconOSS.GetObjectUrl(ctx, avatarKey)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (u *userImpl) ValidateProfileUpdate(ctx context.Context, req *ValidateProfileUpdateRequest) (
	resp *ValidateProfileUpdateResponse, err error) {
	if req.UniqueName == nil && req.Email == nil {
		return nil, fmt.Errorf("missing parameter")
	}

	if req.UniqueName != nil {
		uniqueName := ptr.From(req.UniqueName)
		charNum := utf8.RuneCountInString(uniqueName)

		if charNum < 4 || charNum > 20 {
			return &ValidateProfileUpdateResponse{
				Code: UniqueNameTooShortOrTooLong,
				Msg:  "unique name length should be between 4 and 20",
			}, nil
		}

		exist, err := u.UserRepo.CheckUniqueNameExist(ctx, uniqueName)
		if err != nil {
			return nil, err
		}

		if exist {
			return &ValidateProfileUpdateResponse{
				Code: UniqueNameExist,
				Msg:  "unique name existed",
			}, nil
		}
	}

	return &ValidateProfileUpdateResponse{
		Code: ValidateSuccess,
		Msg:  "success",
	}, nil
}

func (u *userImpl) UpdateProfile(ctx context.Context, req *UpdateProfileRequest) (err error) {
	updates := map[string]interface{}{
		"updated_at": time.Now().UnixMilli(),
	}

	if req.UniqueName != nil {
		resp, err := u.ValidateProfileUpdate(ctx, &ValidateProfileUpdateRequest{
			UniqueName: req.UniqueName,
		})
		if err != nil {
			return err
		}

		if resp.Code != ValidateSuccess {
			return fmt.Errorf("invalid params, %w", err)
		}

		updates["unique_name"] = ptr.From(req.UniqueName)
	}

	if req.Name != nil {
		updates["name"] = ptr.From(req.Name)
	}

	err = u.UserRepo.UpdateProfile(ctx, req.UserID, updates)
	if err != nil {
		return err
	}

	return nil
}

func (u *userImpl) GetUserProfiles(ctx context.Context, userID int64) (user *entity.User, err error) {
	userInfos, err := u.MGetUserProfiles(ctx, []int64{userID})
	if err != nil {
		return nil, err
	}

	if len(userInfos) == 0 {
		return nil, errorx.New(errno.ErrUserNotFoundCode)
	}

	return userInfos[0], nil
}

func (u *userImpl) MGetUserProfiles(ctx context.Context, userIDs []int64) (users []*entity.User, err error) {
	userModels, err := u.UserRepo.GetUsersByIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	users = make([]*userEntity.User, 0, len(userModels))
	for _, um := range userModels {
		// Get image URL
		resURL, err := u.IconOSS.GetObjectUrl(ctx, um.IconURI)
		if err != nil {
			continue // If getting the image URL fails, skip the user
		}

		users = append(users, userPo2Do(um, resURL))
	}

	return users, nil
}

func hashPassword(password string) (string, error) {
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPasswd), nil
}

// Verify that the passwords match
func verifyPassword(password, encodedHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(encodedHash), []byte(password))
}

func userPo2Do(model *model.User, iconURL string) *userEntity.User {
	return &userEntity.User{
		UserID:     model.ID,
		Name:       model.Name,
		UniqueName: model.UniqueName,
		Email:      model.Email,
		IconURI:    model.IconURI,
		IconURL:    iconURL,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}
}
