package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/goalong/bingo/models"
	myError "github.com/goalong/bingo/err"
	"github.com/goalong/bingo/utils"
	"time"
)

func UserRegister(c *gin.Context) {
	var user models.User
	code := myError.SUCCESS
	err := c.BindJSON(&user)
	if err != nil {
		code = myError.INVALID_PARAMS
	}
	filter := make(map[string]interface{})
	filter["email"] = user.Email
	u := models.GetUser(filter)
	if u.ID > 0 {
		code = myError.EMAIL_EXISTS
	} else {
		user.MemberSince = time.Now().Format("2006-01-02 15:04:26")
		models.CreateUser(user)
	}

	c.JSON(
		200,
		gin.H{
			"code": code,
			"msg": myError.GetMsg(code),
		})
}


func Login(c *gin.Context) {
	var user models.User
	filter := make(map[string]interface{})
	err := c.BindJSON(&user)
	tokenString := ""
	code := myError.SUCCESS
	if err != nil {
		code = myError.INVALID_PARAMS
	}
	filter["email"] = user.Email
	realUser := models.GetUser(filter)
	if user.PasswordHash != realUser.PasswordHash {
		code = myError.EMAIL_OR_PW_WRONG
	} else {
		tokenString, err = utils.GenToken(user.Email, user.PasswordHash)
		if err != nil {
			code = myError.SignTokenError
		}

	}
	c.JSON(
		200,
		gin.H{
			"code": code,
			"msg": myError.GetMsg(code),
			"token": tokenString,
		})
}
