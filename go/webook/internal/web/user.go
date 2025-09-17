// user业务层:调用服务层,并返回结果
package web

import (
	"fmt"
	"net/http"
	"time"
	"unicode/utf8"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"github.com/miver02/Learn/go/webook/internal/domain"
	"github.com/miver02/Learn/go/webook/internal/service"
	"github.com/gin-contrib/sessions"
)

// var ErrUserDuplicateEmail = service.ErrUserDuplicateEmail

// user相关路由定义
type UserHandle struct {
	svc 	*service.UserService
	emailExp *regexp.Regexp
	pwdExp *regexp.Regexp
}

func NewUserHandle(svc *service.UserService) *UserHandle {
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	pwdExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandle{
		svc: svc,
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
		ctx.String(http.StatusOK, "参数格式错误")
		fmt.Printf("%v\n", err)
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

	// 调用一下svc的方法
	err = u.svc.SignUp(ctx, domain.User{
		Email: req.Email,
		Password: req.Password,
	})
	if err == service.ErrUserDuplicateEmail {
		ctx.String(http.StatusOK, err.Error())
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		fmt.Printf("%v\n", err)
		return
	}

	ctx.String(http.StatusOK, "注册成功")
	fmt.Printf("%v\n", req)
	// 数据库操作
}

func (u *UserHandle) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email		string `json:"email"`
		Password 	string `json:"password"`
	}

	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	datas_u, err := u.svc.Login(ctx, domain.User{
		Email: 		req.Email,
		Password: 	req.Password,
	})
	if err == service.ErrInvalidUserOrPassword {
		ctx.String(http.StatusOK, err.Error())
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	// 设置session
	sess := sessions.Default(ctx)
	sess.Set("UserId", datas_u.Id)
	sess.Save()

	ctx.String(http.StatusOK, "登录成功")
	fmt.Printf("%v\n", req)
	

}

func (u *UserHandle) Edit(ctx *gin.Context) {
	type EditReq struct {
		Name 			string 	`json:"name"`
		Birthday 		string	`json:"birthday"`
		Introduction 	string 	`json:"introduction"`
	}
	
	var req EditReq
	// 提取数据
	if err := ctx.Bind(&req); err != nil {
		ctx.String(http.StatusOK, "参数格式错误")
		fmt.Printf("%v\n", err)
		return 
	}

	_, err := time.Parse("2006-01-02", req.Birthday)
	if err != nil {
		ctx.String(http.StatusOK, "日期格式不对")
		return
	}

	lenName := utf8.RuneCountInString(req.Name)
	if lenName < 6 && lenName > 15 {
		ctx.String(http.StatusOK, "名字字数不够")
		return
	}

	if utf8.RuneCountInString(req.Introduction) > 20 {
		ctx.String(http.StatusOK, "个人简介内容太多")
		return
	}

	ctx.String(http.StatusOK, "个人信息补充成功")
	fmt.Printf("%v\n", req)
}

func (u *UserHandle) Profile(ctx *gin.Context) {
	ctx.String(http.StatusOK, "这是profile")
}