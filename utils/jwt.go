package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	myError "github.com/goalong/bingo/err"
	"github.com/goalong/bingo/conf"
	"github.com/goalong/bingo/models"

)

type Claims struct {
	Email string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims

}


var jwtSecret = []byte(conf.Config.APP.JwtSecret)


func GenToken(email, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := Claims{
		email,
		password,
		jwt.StandardClaims {
			ExpiresAt : expireTime.Unix(),
			Issuer : "bingo",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}


func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err

}

//检查token的中间件
func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := myError.SUCCESS
		token := c.Query("token")
		var claims *Claims
		var err error
		if token == "" {
			code = myError.INVALID_PARAMS
		} else {
			claims, err = ParseToken(token)
			if err != nil {
				code = myError.CHECK_TOKEN_FAILED
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = myError.CHECK_TOKEN_TIMEOUT
			}
		}
		if code != myError.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : myError.GetMsg(code),
			})

			c.Abort()
			return
		}
		filter := make(map[string]interface{})
		filter["email"] = claims.Email
		user := models.GetUser(filter)
		c.Set("claims", *claims) // 将解析的token数据存入context
		c.Set("user", user)

		c.Next()
	}
}