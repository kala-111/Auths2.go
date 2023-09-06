package controllers

import (
	"github.com/gin-gonic/gin"
	"gin-go-api/config"

)
var :=auths

func GetAuths(c *gin.Context) {
	users := []auths{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}
func CreateAuths(c *gin.Context) {
	var user auths
	c.BindJSON(&user)

	config.DB.Create(&auths)
	c.JSON(200, &user)

}
func DeleteUser(c *gin.Context) {
	var user auths
	config.DB.Where("id=?", c.Param("id")).Delete(&user)
	c.BindJSON(&user)

}
func UpdateUser(c *gin.Context) {
	var user auths
	config.DB.Where("id=?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}
