package server

import (
	"net/http"
	"fmt"
	"time"
	"math"
	
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/spf13/cast"
	jwt "github.com/dgrijalva/jwt-go"

	"local.com/sai0556/Go-ChatwithRabbitMQ/model"
	"local.com/sai0556/Go-ChatwithRabbitMQ/util"
)

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

	token, err := issueToken(user);
	if err != nil {
		SendResponse(c, 1001, "签名生成失败", nil)
		return
	}

	redisClient := model.GetRedis()
	redisClient.HSet(model.Ctx, model.GetKey("register"), token, user.Id)
	SendResponse(c, 0, "登录成功", map[string]string{
		"token": token,
		"id": cast.ToString(user.Id),
		"nickname": user.NickName,
		"avatar": user.Avatar,
	})
}

func Info(c *gin.Context) {
	SendResponse(c, 0, "登录成功", map[string]interface{}{
		"id": c.MustGet("userId"),
		"avatar": c.MustGet("avatar"),
		"nickname": c.MustGet("nickname"),
	})
}

func ChatList(c *gin.Context) {
	db := model.DB
	var list = make([]*model.User, 0)
	var count int
	pageParam := c.DefaultQuery("page", "1")
	limitParam := c.DefaultQuery("limit", "12")
	name := c.DefaultQuery("name", "")
	page := cast.ToInt(pageParam)
	limit := cast.ToInt(limitParam)

	query := db.Where("id != ?", c.MustGet("userId"))
	if name != "" {
		query = query.Where("name like ?", util.StrCombine("%", name, "%"))
	}
	query.Model(&model.User{}).Count(&count)
	if err := query.Select("id, nickname, avatar").Offset(limit*(page - 1)).Limit(limit).Order("id desc").Find(&list).Error;  err != nil {
		SendResponse(c, 3001, "查询失败", nil)
		return 
	}
	SendResponse(c, 0, "ok", map[string]interface{}{
		"data": list,
		"page": page,
		"total": count,
		"limit": limit,
		"total_page": math.Ceil(float64(count)/float64(limit)),
	})
}



type CustomClaims struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Avatar string `json:"avatar"`
	jwt.StandardClaims
}

func issueToken(user model.User) (string,error) {
	mySigningKey := []byte(viper.GetString("token.sign"))

	// Create the Claims
	claims := CustomClaims{
		Id: user.Id,
		Username: user.NickName,
		Avatar: user.Avatar,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("token.ttl")) * time.Minute).Unix(),
			Issuer:    viper.GetString("token.issue"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}


// 解析Token
func ParseToken(tokenString string) (*jwt.Token, *CustomClaims, error) {
    Claims := &CustomClaims{}
    token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
        return  []byte(viper.GetString("token.sign")), nil
    })
    return token, Claims, err
}