package controller

import (
	"api-demo/config"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type Base struct{}

func (t *Base) Api(c *gin.Context, sts int, data interface{}, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": sts,
		"data": data,
		"msg":  msg,
	})
}

func (t *Base) Succ(c *gin.Context, msg string) {
	t.Api(c, config.CODE_SUCCESS, gin.H{}, msg)
}

func (t *Base) Data(c *gin.Context, data interface{}) {
	t.Api(c, config.CODE_SUCCESS, data, "")
}

func (t *Base) Err(c *gin.Context, msg string) {
	t.Api(c, config.CODE_FAIL, gin.H{}, msg)
}

func (t *Base) Paginate(c *gin.Context, db *gorm.DB) error {

	var page, pageSize string
	switch c.Request.Method {
	case "GET":
		page = c.DefaultQuery("page", "1")
		pageSize = c.Query("pageSize")
		break
	case "POST":
		break
	default:
		return errors.New("解析失败")
	}

	var pageN, pageSizeN int
	var err error
	pageN, err = strconv.Atoi(page)
	if err != nil {
		return errors.New("请填写page为数值类型")
	}
	if pageN == 0 {
		pageN = 1
	}
	if pageSize == "" {
		pageSizeN = config.PAGINATE_PAGESIZE
	} else {
		pageSizeN, err = strconv.Atoi(pageSize)
		if err != nil {
			return errors.New("请填写pageSize为数值类型")
		}
		if pageSizeN == 0 {
			pageSizeN = config.PAGINATE_PAGESIZE
		}
	}

	db.Limit(pageSizeN).Offset((pageN - 1) * pageSizeN)
	return nil
}