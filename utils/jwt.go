package utils

import (
	models "hoodhire/structures/models"
	"hoodhire/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   uint
	Username string
	Role     string
	Email    string
	jwt.RegisteredClaims
}

func GenerateTokens(user *models.User)(string,string,error){
	accessTime:=time.Now().Add(time.Hour)
	RefreshTime:=time.Now().Add(7 * 24 * time.Hour)
	accesClaims:=&Claims{
		UserID: user.ID,
		Username: user.Username,
		Email: user.Email,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTime),
			Issuer: "HoodHire",
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Subject: "access-token",
		},
	}
	access:=jwt.NewWithClaims(jwt.SigningMethodHS256,accesClaims)
	accessToken,err:=access.SignedString([]byte(config.AppConfig.JwtKey))
	if err!=nil{
		return "","",err
	}
	refreshClaims:=&Claims{
		UserID: user.ID,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(RefreshTime),
			Issuer: "HoodHire",
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Subject: "refresh-token",
		},
	}
	refresh:=jwt.NewWithClaims(jwt.SigningMethodHS256,refreshClaims)
	refreshToken,err:=refresh.SignedString([]byte(config.AppConfig.JwtKey))
	if err!=nil{
		return "","",err
	}
	return accessToken,refreshToken,nil
}