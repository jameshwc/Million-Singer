package jwt

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jameshwc/Million-Singer/conf"
)

type Claims struct {
	Username string
	Password string
	jwt.StandardClaims
}

const expire = 3 * time.Hour

var jwtSecret []byte

func init() {
	jwtSecret = []byte(conf.ServerConfig.JwtSecret)
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(expire)

	claims := Claims{
		encodeMD5(username),
		encodeMD5(password),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "million-singer",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func encodeMD5(s string) string {
	m := md5.New()
	m.Write([]byte(s))

	return hex.EncodeToString(m.Sum(nil))
}
