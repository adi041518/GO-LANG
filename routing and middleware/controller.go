package controllers

import (
	"net/http"

	"example1.com/user-registartion-api/model"
	"example1.com/user-registartion-api/services"
	"github.com/gin-gonic/gin"
)

type Usercontroller struct {
	ser *services.Users
}

func NewController(ser *services.Users) *Usercontroller {
	return &Usercontroller{ser: ser}
}

func (u *Usercontroller) UserValidateController(c *gin.Context) {
	var m model.User

	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	res := u.ser.Uservalidate(m)

	c.JSON(http.StatusOK, gin.H{"result": res})
}
