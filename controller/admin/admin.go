package admin

import (
	"api-demo/controller"
	"api-demo/lib/database"
	"api-demo/model"
	"github.com/gin-gonic/gin"
	"github.com/zctod/tool/common/utils"
	"strconv"
)

type Admin struct{ controller.Base }

// 后台首页
func (t *Admin) Index(c *gin.Context) {
	t.Succ(c, "Hello World!")
}

// 管理员列表
func (t *Admin) List(c *gin.Context) {

	db, err := database.Open()
	defer db.Close()
	if err != nil {
		t.Err(c, "系统错误")
		return
	}

	var admins []model.Admin
	err = t.Paginate(c, db)
	if err != nil {
		t.Err(c, err.Error())
		return
	}
	db.Find(&admins)

	t.Data(c, utils.JsonToMap(admins))
}

// 管理员添加
func (t *Admin) Create(c *gin.Context) {

	_ = c.Request.ParseForm()
	name := c.PostForm("name")
	password := c.PostForm("password")
	roleId := c.PostForm("roleId")

	if name == "" {
		t.Err(c, "请输入管理员名称")
		return
	}
	if password == "" {
		t.Err(c, "请输入管理员密码")
		return
	}
	if roleId == "" {
		t.Err(c, "请选择管理员分组")
		return
	}

	db, err := database.Open()
	defer db.Close()
	if err != nil {
		t.Err(c, "系统错误")
		return
	}

	adminRole := model.AdminRole{}
	db.Where("id = ?", roleId).First(&adminRole)
	if adminRole.ID == 0 {
		t.Err(c, "不存在的管理员分组")
		return
	}

	admin := model.Admin{
		Name: name,
	}
	db.First(&admin, admin)
	if admin.ID == 0 {
		admin.Password = utils.MD5(password)
		admin.Role = adminRole.ID
		db.Create(&admin)
		if db.NewRecord(admin) == true {
			t.Err(c, "添加失败")
			return
		}
	}
	t.Succ(c, "添加成功")
}

// 管理员修改
func (t *Admin) Update(c *gin.Context) {

	_ = c.Request.ParseForm()
	id := c.DefaultPostForm("id", "0")
	name := c.PostForm("name")
	password := c.PostForm("password")
	roleId := c.PostForm("roleId")
	if id == "0" {
		t.Err(c, "请选择管理员")
		return
	}
	if name == "" && password == "" && roleId == "" {
		t.Err(c, "无修改内容")
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
	if name != "" {
		adminSelect := model.Admin{
			Name: name,
		}
		db.First(&adminSelect, adminSelect)
		if adminSelect.ID > 0 {
			t.Err(c, "该管理员名称已被占用")
			return
		}
		admin.Name = name
	}
	if password != "" {
		admin.Password = utils.MD5(password)
	}
	roleIdN, _ := strconv.Atoi(roleId)
	if roleIdN > 0 {
		adminRole := model.AdminRole{}
		db.Where("id = ?", roleIdN).First(&adminRole)
		if adminRole.ID == 0 {
			t.Err(c, "不存在的管理员分组")
			return
		}
		admin.Role = uint(roleIdN)
	}
	db.Save(&admin)

	t.Succ(c, "修改成功")
}

// 管理员删除
func (t *Admin) Delete(c *gin.Context) {

	_ = c.Request.ParseForm()
	id := c.PostForm("id")

	db, err := database.Open()
	defer db.Close()
	if err != nil {
		t.Err(c, "系统错误")
		return
	}

	admin := model.Admin{}
	db.Where("id = ?", id).First(&admin)
	if admin.ID == 0 {
		t.Err(c, "不存在的管理员")
		return
	}

	db.Delete(&admin)
	t.Succ(c, "删除成功")
}

// 管理员分组列表
func (t *Admin) RoleList(c *gin.Context) {

	db, err := database.Open()
	defer db.Close()
	if err != nil {
		t.Err(c, "系统错误")
		return
	}

	var roles []model.AdminRole
	err = t.Paginate(c, db)
	if err != nil {
		t.Err(c, err.Error())
		return
	}
	db.Find(&roles)

	t.Data(c, utils.JsonToMap(roles))
}

// 管理员分组添加
func (t *Admin) RoleCreate(c *gin.Context) {

	_ = c.Request.ParseForm()
	name := c.PostForm("name")
	rule := c.PostForm("rule")
	if name == "" {
		t.Err(c, "请填写管理员分组名称")
		return
	}

	db, err := database.Open()
	defer db.Close()
	if err != nil {
		t.Err(c, "系统错误")
		return
	}
	role := model.AdminRole{
		Name: name,
	}
	db.First(&role, role)
	if role.ID == 0 {
		role.Rule = rule
		db.Create(&role)
		if db.NewRecord(role) == true {
			t.Err(c, "添加失败")
			return
		}
	}
	t.Succ(c, "添加成功")
}

// 管理员分组修改
func (t *Admin) RoleUpdate(c *gin.Context) {

	_ = c.Request.ParseForm()
	id := c.PostForm("id")
	name := c.PostForm("name")
	rule := c.PostForm("rule")

	db, err := database.Open()
	defer db.Close()
	if err != nil {
		t.Err(c, "系统错误")
		return
	}

	role := model.AdminRole{}
	db.Where("id = ?", id).First(&role)
	if role.ID == 0 {
		t.Err(c, "不存在的管理员分组")
		return
	}
	if name != "" {
		roleSelect := model.AdminRole{}
		db.Where("name = ?", name).First(&roleSelect)
		if roleSelect.ID > 0 {
			t.Err(c, "管理员分组名称已存在")
			return
		}
		role.Name = name
	}
	role.Rule = rule
	db.Save(&role)
	t.Succ(c, "操作成功")
}

// 管理员分组删除
func (t *Admin) RoleDelete(c *gin.Context) {

	_ = c.Request.ParseForm()
	id := c.PostForm("id")

	db, err := database.Open()
	defer db.Close()
	if err != nil {
		t.Err(c, "系统错误")
	}

	role := model.AdminRole{}
	db.Where("id = ?", id).First(&role)
	if role.ID == 0 {
		t.Err(c, "不存在的管理员分组")
		return
	}

	idN, _ := strconv.Atoi(id)
	admin := model.Admin{
		Role: uint(idN),
	}
	db.First(&admin, admin)
	if admin.ID > 0 {
		t.Err(c, "管理员分组还存在管理员，无法删除")
		return
	}

	db.Delete(&role)
	t.Succ(c, "操作成功")
}
