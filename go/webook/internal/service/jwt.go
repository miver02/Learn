package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/miver02/learn-program/go/webook/internal/consts"
)

type JwtService struct {
}

func NewJwtServer() *JwtService {
	return &JwtService{
	}
}

func (js *JwtService) SetJwtToken(ctx *gin.Context, id int64) error {
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 3600 * 24)),
		},
		Uid:       id,
		UserAgent: ctx.Request.UserAgent(),
	}
	// 使你的token携带参数
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString([]byte("secret"))
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return err
	}
	ctx.Header("x-jwt-token", "Bearer "+tokenStr)
	return nil
}

func (js *JwtService) GetIdByJwtClaims(ctx *gin.Context) (int64, error) {
	// 获取jwt的claims
	claimsVal, ok := ctx.Get("claims")
	if !ok {
		ctx.String(http.StatusUnauthorized, "未登录或会话失效")
		return 0, consts.ErrJwtTokenInvalid
	}
	claims, ok := claimsVal.(*UserClaims)
	if !ok {
		ctx.String(http.StatusUnauthorized, "系统错误")
		return 0, consts.ErrSystem
	}

	return claims.Uid, nil
}

type UserClaims struct {
	// token携带参数接口
	jwt.RegisteredClaims
	Uid       int64
	UserAgent string
}