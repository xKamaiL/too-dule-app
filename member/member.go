package member

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/xkamail/too-dule-app/pkg/config"
	"github.com/xkamail/too-dule-app/pkg/logger"
	"time"
)

func CreateToken(userID string) (string, error) {
	cfg := config.Load()
	var err error
	// Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(cfg.JWTSecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

type Service struct {
	repo Repository
	cfg  *config.Config
	l    *logger.Logger
}

func New() *Service {
	return &Service{
		repo: Repository{},
		cfg:  config.Load(),
		l:    logger.New("Member"),
	}
}

type CreateMemberParam struct {
	Username string `json:"username" validate:"required,min=4"`
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
}

func (s *Service) Create(ctx context.Context, params CreateMemberParam) (string, error) {
	userId, err := s.repo.Insert(ctx, struct {
		Username string
		Password string
		Email    string
	}(params))
	if err != nil {
		s.l.Info(err)
		return "", err
	}
	return CreateToken(userId)
}
