package dal

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/crazyfrankie/ddd-todolist/backend/domain/user/internal/dal/model"
	"github.com/crazyfrankie/ddd-todolist/backend/domain/user/internal/dal/query"
)

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		query: query.Use(db),
	}
}

type UserDAO struct {
	query *query.Query
}

func (dao *UserDAO) GetUsersByEmail(ctx context.Context, email string) (*model.User, bool, error) {
	user, err := dao.query.User.WithContext(ctx).Where(dao.query.User.Email.Eq(email)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, false, nil
	}

	if err != nil {
		return nil, false, err
	}

	return user, true, err
}

func (dao *UserDAO) UpdateSessionKey(ctx context.Context, userID int64, sessionKey string) error {
	_, err := dao.query.User.WithContext(ctx).Where(
		dao.query.User.ID.Eq(userID),
	).Updates(map[string]any{
		"session_key": sessionKey,
		"updated_at":  time.Now().UnixMilli(),
	})
	return err
}

func (dao *UserDAO) ClearSessionKey(ctx context.Context, userID int64) error {
	_, err := dao.query.User.WithContext(ctx).
		Where(
			dao.query.User.ID.Eq(userID),
		).
		UpdateColumn(dao.query.User.SessionKey, "")

	return err
}

func (dao *UserDAO) UpdatePassword(ctx context.Context, email, password string) error {
	_, err := dao.query.User.WithContext(ctx).Where(
		dao.query.User.Email.Eq(email),
	).Updates(map[string]any{
		"password":    password,
		"session_key": "", // clear session key
		"updated_at":  time.Now().UnixMilli(),
	})
	return err
}

func (dao *UserDAO) GetUserByID(ctx context.Context, userID int64) (*model.User, error) {
	return dao.query.User.WithContext(ctx).Where(
		dao.query.User.ID.Eq(userID),
	).First()
}

func (dao *UserDAO) UpdateAvatar(ctx context.Context, userID int64, iconURI string) error {
	_, err := dao.query.User.WithContext(ctx).Where(
		dao.query.User.ID.Eq(userID),
	).Updates(map[string]any{
		"icon_uri":   iconURI,
		"updated_at": time.Now().UnixMilli(),
	})
	return err
}

func (dao *UserDAO) CheckUniqueNameExist(ctx context.Context, uniqueName string) (bool, error) {
	_, err := dao.query.User.WithContext(ctx).Select(dao.query.User.ID).Where(
		dao.query.User.UniqueName.Eq(uniqueName),
	).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (dao *UserDAO) UpdateProfile(ctx context.Context, userID int64, updates map[string]any) error {
	if _, ok := updates["updated_at"]; !ok {
		updates["updated_at"] = time.Now().UnixMilli()
	}

	_, err := dao.query.User.WithContext(ctx).Where(
		dao.query.User.ID.Eq(userID),
	).Updates(updates)
	return err
}

func (dao *UserDAO) CheckEmailExist(ctx context.Context, email string) (bool, error) {
	_, exist, err := dao.GetUsersByEmail(ctx, email)
	if !exist {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

// CreateUser Create a new user
func (dao *UserDAO) CreateUser(ctx context.Context, user *model.User) error {
	return dao.query.User.WithContext(ctx).Create(user)
}

// GetUserBySessionKey Query users based on session key
func (dao *UserDAO) GetUserBySessionKey(ctx context.Context, sessionKey string) (*model.User, bool, error) {
	sm, err := dao.query.User.WithContext(ctx).Where(
		dao.query.User.SessionKey.Eq(sessionKey),
	).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}

	return sm, true, nil
}

// GetUsersByIDs Query user information in batches
func (dao *UserDAO) GetUsersByIDs(ctx context.Context, userIDs []int64) ([]*model.User, error) {
	return dao.query.User.WithContext(ctx).Where(
		dao.query.User.ID.In(userIDs...),
	).Find()
}
