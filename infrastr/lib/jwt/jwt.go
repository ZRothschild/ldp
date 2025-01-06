package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"time"
)

type (
	MyClaims struct {
		jwt.RegisteredClaims
	}
)

// SignJwt 生成jwt
func SignJwt(sk, nickname string) (t string, err error) {
	var (
		mapClaims = MyClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "ldp发型",
				Subject:   nickname,
				Audience:  jwt.ClaimStrings{"受众为前段资源"},
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	)
	if t, err = token.SignedString(sk); err != nil {
		return t, err
	}
	return t, err
}

// ParseJwt 解析jwt
func ParseJwt(sk any, jwtStr string, options ...jwt.ParserOption) (jwt.Claims, error) {
	var (
		mapClaims = new(MyClaims)
	)
	token, err := jwt.ParseWithClaims(jwtStr, mapClaims, func(token *jwt.Token) (interface{}, error) {
		return sk, nil
	}, options...)
	if err != nil {
		return nil, err
	}
	// 校验 Claims 对象是否有效，基于 exp（过期时间），nbf（不早于），iat（签发时间）等进行判断（如果有这些声明的话）。
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token.Claims, nil
}
