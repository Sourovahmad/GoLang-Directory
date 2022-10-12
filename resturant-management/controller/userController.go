package controller

import (
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	//
}



func GetUser(c *gin.Context) {
	//
}




func SingUp(c *gin.Context) {
	//
}


func Login(c *gin.Context) {
	//
}


func HashPassword(password string) string{
	return "hello"
}

func VarifyPassword(userPassword string,proidePassword string)(bool, string){
	//
	return true, "hello"
}
