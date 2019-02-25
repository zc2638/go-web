package admin

import (
	"api-demo/config"
	"api-demo/controller"
	"api-demo/lib/database"
	"api-demo/lib/jwt"
	"api-demo/model"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
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

	var m = md5.New()
	m.Write([]byte(password))

	var admin = model.Admin{
		Name:     name,
		Password: hex.EncodeToString(m.Sum(nil)),
	}
	db.First(&admin, admin)
	if admin.ID == 0 {
		t.Err(c, "不存在的管理员")
		return
	}

	var data = map[string]interface{}{
		"id": admin.ID,
		"name": admin.Name,
		"role": admin.Role,
	}
	token, err := jwt.Create(data, config.JWT_SECRET_ADMIN, time.Now().Add(time.Hour * config.JWT_EXP_ADMIN).Unix())
	if err != nil {
		t.Err(c, "登陆失败")
		return
	}
	t.Data(c, token)
}

// 登出
func (t *Auth) Logout(c *gin.Context) {

	t.Succ(c, "登出成功")
}
