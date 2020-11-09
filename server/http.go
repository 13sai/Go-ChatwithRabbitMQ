package server

import (
	"net/http"
	"fmt"
	"strings"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	jwt "github.com/dgrijalva/jwt-go"

	"local.com/sai0556/Go-ChatwithRabbitMQ/model"
	"local.com/sai0556/Go-ChatwithRabbitMQ/util"
)

func Login(c *gin.Context) {
	type Params struct {
		Username string `json:"username" form:"username" binding:"required"`
		Password  string `json:"password" form:"password" binding:"required"`
	}

	form := Params{}
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(form)
		SendResponse(c, 1001, "参数错误", nil)
		return
	}

	var user = model.User{}
	db := model.DB
	db.Where("nickname = ?", form.Username).First(&user)

	if user.Id < 1 {
		fmt.Println(user)
		SendResponse(c, 2001, "昵称或密码不匹配", nil)
		return
	}

	if user.Password != util.Md5(util.StrCombine(viper.GetString("sign"), form.Password)) {
		fmt.Println(user)
		SendResponse(c, 2002, "昵称或密码不匹配", nil)
		return
	}

	token, err := issueToken(form.Username, form.Password);
	if err != nil {
		SendResponse(c, 1001, "签名生成失败", nil)
		return
	}

	redisClient := model.GetRedis()
	redisClient.HSet(model.Ctx, model.GetKey("register"), token, user.Id)
	SendResponse(c, 0, "登录成功", token)
}

func Info(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")

	if len(auth) == 0 {
		SendResponse(c, http.StatusUnauthorized, "未登录", nil)
		return 
	}

	auth = strings.Fields(auth)[1]
	// 校验token
	_, Claims, err := parseToken(auth)
	if err != nil {
		fmt.Println(err)
		SendResponse(c, http.StatusUnauthorized, "未登录", nil)
		return 
	} else {
		fmt.Println("token 正确")
	}

	SendResponse(c, 0, "登录成功", Claims)
}

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func SendResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Message: message,
		Data: data,
	})
}

type CustomClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func issueToken(username ,password string) (string,error) {
	mySigningKey := []byte(viper.GetString("token.sign"))

	// Create the Claims
	claims := CustomClaims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("token.ttl")) * time.Minute).Unix(),
			Issuer:    viper.GetString("token.issue"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}


// 解析Token
func parseToken(tokenString string) (*jwt.Token, *CustomClaims, error) {
    Claims := &CustomClaims{}
    token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
        return  []byte(viper.GetString("token.sign")), nil
    })
    return token, Claims, err
}