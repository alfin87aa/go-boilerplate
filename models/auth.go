package models

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

//TokenDetails ...
type TokenDetails struct {
	AccessToken         string
	RefreshToken        string
	AccessUUID          string
	RefreshUUID         string
	AcessTokenExpires   time.Time
	RefreshTokenExpires time.Time
}

//CreateToken ...
func (entity *TokenDetails) CreateToken(userID string) (*TokenDetails, error) {

	entity.AcessTokenExpires = time.Now().Add(time.Minute * 15)
	entity.AccessUUID = uuid.New().String()

	entity.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7)
	entity.RefreshUUID = uuid.New().String()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = entity.AccessUUID
	atClaims["user_id"] = userID
	atClaims["exp"] = entity.AcessTokenExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	entity.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = entity.RefreshUUID
	rtClaims["user_id"] = userID
	rtClaims["exp"] = entity.RefreshTokenExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	entity.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return entity, nil
}
