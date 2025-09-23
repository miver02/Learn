// user业务层:调用服务层,并返回结果
package web

import (
	"fmt"
	"net/http"
	"time"
	"unicode/utf8"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/miver02/Learn/go/webook/internal/domain"
	"github.com/miver02/Learn/go/webook/internal/service"
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
		Email: 		req.Email,
		Password: 	req.Password,
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

/*
func (u *UserHandle) LoginSession(ctx *gin.Context) {
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
	sess.Options(sessions.Options{
		MaxAge:   10 * 60, // 10分钟，按你的需求设置
		Path:     "/",     // 确保路径匹配，避免多个 Cookie 冲突
		Domain:   "127.0.0.1", // 明确域名，与请求一致
		HttpOnly: true,
		Secure:   false,   // 本地开发关闭 Secure
	})
	sess.Save()

	ctx.String(http.StatusOK, "登录成功")
	fmt.Printf("%v\n", req)
}
*/

func (u *UserHandle) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email		string `json:"email"`
		Password 	string `json:"password"`
	}

	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	_, err := u.svc.Login(ctx, domain.User{
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

	// JWT
	token := jwt.New(jwt.SigningMethodHS512)
	tokenStr, err := token.SignedString([]byte("secret"))
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	println(tokenStr)
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
	
	sess := sessions.Default(ctx)
	idVal := sess.Get("UserId")
	id, ok := idVal.(int64)
	if !ok {
		ctx.String(http.StatusUnauthorized, "未登录或会话失效")
		return
	}
	err = u.svc.Edit(ctx, domain.User{
		Id:           id,
		Name:         req.Name,
		Birthday:     req.Birthday,
		Introduction: req.Introduction,
	})
	
	if err != nil {
		ctx.String(http.StatusUnauthorized, "个人信息补充失败")
		return
	}

	ctx.String(http.StatusOK, "个人信息补充成功")
	fmt.Printf("%v\n", req)
}

func (u *UserHandle) Profile(ctx *gin.Context) {
	ctx.String(http.StatusOK, "这是profile")
}

func (u *UserHandle) Logout(ctx *gin.Context) {
	// 1. 获取当前会话
	sess := sessions.Default(ctx)
	
	// 2. 清除会话中的用户信息（关键步骤）
	sess.Delete("UserId") // 删除存储的用户ID
	
	// 可选：设置会话立即过期（彻底销毁会话）
	// sess.Options(sessions.Options{
	// 	MaxAge: -1, // MaxAge=-1 表示立即过期
	// })

	// 3. 保存会话状态（必须调用，否则删除操作不生效）
	if err := sess.Save(); err != nil {
		ctx.String(http.StatusOK, "退出登录失败")
		return
	}

	ctx.String(http.StatusOK, "退出登录成功")
	fmt.Printf("UserId: %d\n", sess.Get("UserId"))
}