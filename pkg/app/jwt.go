package app

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/nico612/blog-service/global"
	"github.com/nico612/blog-service/pkg/util"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.RegisteredClaims
}

// GetJWTSecret 获取JWT的 secret
func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// GenerateToken 生成token
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire) //过期时间
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	// 根据指定的算法生成签名后的token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成token string
	return tokenClaims.SignedString(GetJWTSecret())
}

// ParseToken 解析Token
// 它主要的功能是解析和校验 Token
func ParseToken(token string) (*Claims, error) {

	// 解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回 *Token。
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		// Valid：验证基于时间的声明，例如：过期时间（ExpiresAt）、签发者（Issuer）、生效时间（Not Before），需要注意的是，如果没有任何声明在令牌中，仍然会被认为是有效的。
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
