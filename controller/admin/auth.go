package admin

import (
	"go-web/config"
	"go-web/controller"
	"go-web/lib/database"
	"go-web/lib/jwt"
	"go-web/model"
	"github.com/gin-gonic/gin"
	"github.com/zctod/tool/common/utils"
	"time"
)

type Auth struct{ controller.Base }

// 登陆
func (t *Auth) Login(c *gin.Context) {

	_ = c.Request.ParseForm()

	name := c.PostForm("name")
	password := c.PostForm("password")
	if name == "" {
		t.Err(c, "请输入用户名")
		return
	}
	if password == "" {
		t.Err(c, "请输入用户名密码")
		return
	}

	db, err := database.Open()
	defer db.Close()
	if err != nil {
		t.Err(c, "系统错误")
		return
	}

	var admin = model.Admin{
		Name:     name,
		Password: utils.MD5(password),
	}
	db.First(&admin, admin)
	if admin.ID == 0 {
		t.Err(c, "管理员账号密码错误")
		return
	}

	var adminRole = model.AdminRole{}
	db.Where("id = ?", admin.Role).First(&adminRole)
	if adminRole.ID == 0 {
		t.Err(c, "当前管理员所在分组不存在")
		return
	}

	var data = map[string]interface{}{
		"id":   admin.ID,
		"name": admin.Name,
		"role": admin.Role,
	}
	token, err := jwt.Create(data, config.JWT_SECRET_ADMIN, time.Now().Add(time.Hour * config.JWT_EXP_ADMIN).Unix())
	if err != nil {
		t.Err(c, "登陆失败")
		return
	}
	t.Data(c, gin.H{
		"token": token,
		"name":  admin.Name,
		"rule":  adminRole.Rule,
	})
}

// 登出
func (t *Auth) Logout(c *gin.Context) {

	t.Succ(c, "登出成功")
}

// 个人详情
func (t *Auth) Show(c *gin.Context) {

	tokenStr := c.Request.Header.Get("token")
	if tokenStr == "" {
		t.Err(c, "请先登录")
		return
	}
	jwtData, err := jwt.ParseInfo(tokenStr, config.JWT_SECRET_ADMIN)
	if err != nil {
		t.Err(c, "异常登录信息1")
		return
	}
	info , ok := jwtData["info"]
	if !ok {
		t.Err(c, "异常登录信息2")
		return
	}
	id, ok := info.(map[string]interface{})["id"]
	if !ok {
		t.Err(c, "异常登录信息3")
		return
	}
	roleId, ok := info.(map[string]interface{})["role"]
	if !ok {
		t.Err(c, "异常登录信息4")
		return
	}

	db, err := database.Open()
	defer db.Close()
	if err != nil {
		t.Err(c, "系统错误")
		return
	}

	var admin = model.Admin{}
	db.Where("id = ?", id).First(&admin)
	if admin.ID == 0 {
		t.Err(c, "不存在的管理员")
		return
	}

	var adminRole = model.AdminRole{}
	db.Where("id = ?", roleId).First(&adminRole)
	if adminRole.ID == 0 {
		t.Err(c, "不存在的管理员分组")
		return
	}

	var data = map[string]interface{}{
		"id":   admin.ID,
		"name": admin.Name,
		"role": admin.Role,
	}
	token, err := jwt.Create(data, config.JWT_SECRET_ADMIN, time.Now().Add(time.Hour * config.JWT_EXP_ADMIN).Unix())
	if err != nil {
		t.Err(c, "操作失败")
		return
	}

	t.Data(c, gin.H{
		"token": token,
		"name": admin.Name,
		"rule": adminRole.Rule,
	})
}
