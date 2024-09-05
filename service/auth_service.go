package service

import (
	"context"
	"errors"
	"time"

	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/config"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/model"
	errpkg "github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/pkg/error"
	jwtpkg "github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/pkg/jwt"
	md5pkg "github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/pkg/md5"
	timepkg "github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/pkg/time"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/repository"
)

type AuthService interface {
	Login(ctx context.Context, username, password string) (*string, error)
	Register(ctx context.Context, user model.User) (*model.User, error)
}

type authServiceImpl struct {
	config   *config.Config
	authRepo repository.AuthRepository
}

func NewAuthService(
	config *config.Config,
	authRepo repository.AuthRepository,
) AuthService {
	return &authServiceImpl{
		config:   config,
		authRepo: authRepo,
	}
}

func (s *authServiceImpl) Login(ctx context.Context, username, password string) (*string, error) {
	user, err := s.authRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New(errpkg.GetMsg(errpkg.ERROR_AUTH_USER_NOT_EXIST))
	}

	if ok := md5pkg.CompareMD5(password, user.Password); !ok {
		return nil, errors.New(errpkg.GetMsg(errpkg.ERROR_AUTH_PASSWORD_INCORRECT))
	}

	token, err := s.generateToken(*user)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *authServiceImpl) Register(ctx context.Context, user model.User) (*model.User, error) {
	userExist, err := s.authRepo.FindByUsername(ctx, user.Username)
	if err != nil && err.Error() != "mongo: no documents in result" {
		return nil, err
	}

	if userExist != nil {
		return nil, errors.New(errpkg.GetMsg(errpkg.ERROR_AUTH_USERNAME_EXIST))
	}

	user.Password = md5pkg.EncodeMD5(user.Password)

	newUser, err := s.authRepo.Create(ctx, &user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *authServiceImpl) generateToken(user model.User) (*string, error) {
	expiration := timepkg.StrToDuration(s.config.JWT.Expiration)
	payload := map[string]interface{}{
		"username": user.Username,
		"sub":      user.ID.Hex(),
		"exp":      time.Now().Add(expiration).Unix(),
	}
	token, err := jwtpkg.GenerateToken(payload, s.config.JWT.Secret)

	if err != nil {
		return nil, err
	}

	return &token, nil
}
