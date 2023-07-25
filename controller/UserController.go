package controller

import (
	"ginvue/common"
	"ginvue/model"
	"ginvue/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
)

func Register(c *gin.Context) {
	db := common.GetDB()
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	if len(telephone) != 11 {
		c.JSON(422, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		c.JSON(422, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Print(name, telephone, password)
	//判断手机号是否存在
	if isTelephoneExist(db, telephone) {
		c.JSON(422, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}
	//创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	db.Create(&newUser)
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

//查询手机号

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false

}
