package member

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/xkamail/too-dule-app/pkg/config"
	"time"
)

func CreateToken(userid string) (string, error) {
	cfg := config.Load()
	var err error
	// Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(cfg.JWTSecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
