package web

import (
	"fmt"
	"net/http"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

// user相关路由定义
type UserHandle struct {
	emailExp *regexp.Regexp
	pwdExp *regexp.Regexp
}

func NewUserHandle() *UserHandle {
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	pwdExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandle{
		emailExp: emailExp,
		pwdExp: pwdExp,
	}
}


func (u *UserHandle) SignUp(ctx *gin.Context) {
	type SignUpReq struct{
		// 内部结构体
		Email 			string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password 		string `json:"password"`
	}

	var req SignUpReq
	// Bind会根据Content-Type类型解析你的数据到req中
	// 解析错了会返回400错误
	if err := ctx.Bind(&req); err != nil {
		return
	}
	// 邮箱效验
	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "邮箱格式错误")
		return
	}


	// 密码效验
	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusOK, "两次输入的密码不一致")
		return
	}
	ok, err = u.pwdExp.MatchString(req.Password)
	if err != nil {
		// 记录日志
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	if !ok {
		ctx.String(http.StatusOK, "密码必须大于八位,且包含字母和特殊字符")
		return
	}


	ctx.String(http.StatusOK, "注册成功")
	fmt.Printf("%v\n", req)
	// 数据库操作
}

func (u *UserHandle) Login(ctx *gin.Context) {
	
	
}

func (u *UserHandle) Edit(ctx *gin.Context) {
	
	
}

func (u *UserHandle) Profile(ctx *gin.Context) {
	
	
}