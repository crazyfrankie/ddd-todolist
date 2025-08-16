package token

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"

	"github.com/crazyfrankie/ddd-todolist/backend/infra/contract/token"
)

type JWT = token.JWT

type jwtImpl struct {
	cmd       redis.Cmdable
	signAlgo  string
	secretKey []byte
}

func New(cmd redis.Cmdable, signAlgo string, secret string) token.JWT {
	return &jwtImpl{cmd: cmd, signAlgo: signAlgo, secretKey: []byte(secret)}
}

func (s *jwtImpl) GenerateToken(uid int64, ua string) ([]string, error) {
	res := make([]string, 2)
	access, err := s.newToken(uid, time.Minute*15)
	if err != nil {
		return res, err
	}
	res[0] = access
	refresh, err := s.newToken(uid, time.Hour*24*30)
	if err != nil {
		return res, err
	}
	res[1] = refresh

	// set refresh in redis
	key := tokenKey(uid, ua)

	err = s.cmd.Set(context.Background(), key, refresh, time.Hour*24*30).Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *jwtImpl) newToken(uid int64, duration time.Duration) (string, error) {
	now := time.Now()
	claims := &token.Claims{
		UserID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
		},
	}
	newToken := jwt.NewWithClaims(jwt.GetSigningMethod(s.signAlgo), claims)
	str, err := newToken.SignedString(s.secretKey)

	return str, err
}

func (s *jwtImpl) ParseToken(tk string) (*token.Claims, error) {
	t, err := jwt.ParseWithClaims(tk, &token.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(*token.Claims)
	if ok {
		return claims, nil
	}

	return nil, errors.New("jwt is invalid")
}

func (s *jwtImpl) TryRefresh(refresh string, ua string) ([]string, *token.Claims, error) {
	refreshClaims, err := s.ParseToken(refresh)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid refresh jwt")
	}

	res, err := s.cmd.Get(context.Background(), tokenKey(refreshClaims.UserID, ua)).Result()
	if err != nil || res != refresh {
		return nil, nil, errors.New("jwt invalid or revoked")
	}

	access, err := s.newToken(refreshClaims.UserID, time.Hour)
	if err != nil {
		return nil, nil, err
	}

	now := time.Now()
	issat, _ := refreshClaims.GetIssuedAt()
	expire, _ := refreshClaims.GetExpirationTime()
	if expire.Sub(now) < expire.Sub(issat.Time)/3 {
		// try refresh
		refresh, err = s.newToken(refreshClaims.UserID, time.Hour*24*30)
		err = s.cmd.Set(context.Background(), tokenKey(refreshClaims.UserID, ua), refresh, time.Hour*24*30).Err()
		if err != nil {
			return nil, nil, err
		}
	}

	return []string{access, refresh}, refreshClaims, nil
}

func (s *jwtImpl) CleanToken(ctx context.Context, uid int64, ua string) error {
	return s.cmd.Del(ctx, tokenKey(uid, ua)).Err()
}

func (s *jwtImpl) GetAccessToken(c *gin.Context) (string, error) {
	tokenHeader := c.GetHeader("Authorization")
	if tokenHeader == "" {
		return "", errors.New("no auth")
	}

	strs := strings.Split(tokenHeader, " ")
	if len(strs) != 2 || strs[0] != "Bearer" {
		return "", errors.New("header is invalid")
	}

	if strs[1] == "" {
		return "", errors.New("jwt is empty")
	}

	return strs[1], nil
}

func tokenKey(uid int64, ua string) string {
	hash := hashUA(ua)
	return fmt.Sprintf("refresh_token:%d:%s", uid, hash)
}

func hashUA(ua string) string {
	sum := sha1.Sum([]byte(ua))
	return hex.EncodeToString(sum[:])
}
