package middleware

import (
	"fmt"
	"net/http"
	"spm-serv/model/po"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)


//盐值
const (
	// 可自定义盐值
	key = "toekn_key"
)

type LoginRsp struct {
	Token string `json:"token"`
}

func verifyAction(c *gin.Context) (bool){
	var b bool = true
	token := c.Request.Header.Get("token")
	if token == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusUnauthorized,
			"msg":    "请求未携带token，无权限访问",
		})
		//c.Abort()
		return false
	}
	return b
}

//
func Authorize() gin.HandlerFunc{
	return func(c *gin.Context){
		// 获取访问令牌
		tokenString := c.Request.Header.Get("token")
		if tokenString == "" {
			Abort(c, "请求未携带token，无权限访问")
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})

		if token != nil && token.Valid {
			fmt.Println("You look nice today")
			//正常
			c.Next()
			// 继续交由下一个路由处理,并将解析出的信息传递下去
			//验证通过执行正常流程代码
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			//验证失败，返回提示
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				Abort(c, "这不是一个正确的token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				//过期
				Abort(c, "token已过期")
			} else {
				Abort(c, "这不是一个正确的token")
			}
		} else {
			Abort(c, "这不是一个正确的token")
		}



	}
}

//错误处理
func Abort(c *gin.Context, error string){
	c.Abort()
	//c.JSON(http.StatusUnauthorized, gin.H{"Msg":error})//token无效
	// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusUnauthorized,
		"msg":  error,
	})
	return
}

/*

// 一些常量
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "newtrekWang"
)
*/
//生成token
func SetToken(consoleUser po.ConsoleUser) string {
	//1。生成token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	//2。添加令牌关键信息
	//添加令牌期限
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["id"]=consoleUser.Id
	claims["userName"]=consoleUser.Username
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return ""
	}
	return tokenString
}



