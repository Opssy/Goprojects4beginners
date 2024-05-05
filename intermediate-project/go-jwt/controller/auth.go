package controller

import (
	"go-awt/model"
	"go-awt/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_secret_key")


func Login(c *gin.Context){

	var user model.User

	if err := c.ShouldBindJSON(&user);
	err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
    
	 var  existingUser model.User

	 model.DB.Where("email = ?", user.Email).First(&existingUser)

	 if existingUser.ID == 0 {
		    c.JSON(400, gin.H{"error": "user does not exist"})
          return
	 }

	 errHash := utils.CompareHashPassword(user.Password, existingUser.Password)

	 if !errHash{
		 c.JSON(400, gin.H{"error": "invalid password"})
          return
	 }
	 expirationTime := time.Now().Add(5 * time.Minute)

	 claims := model.Claims{
		   Role: existingUser.Role,
          StandardClaims: jwt.StandardClaims{
              Subject:   existingUser.Email,
              ExpiresAt: expirationTime.Unix(),
          },
	 }
	     token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

      tokenString, err := token.SignedString(jwtKey)

      if err != nil {
          c.JSON(500, gin.H{"error": "could not generate token"})
          return
      }
	  c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
      c.JSON(200, gin.H{"success": "user logged in"})
}


func Signup(c *gin.Context){
	var user model.User


	if err := c.ShouldBindJSON(&user); 

	err != nil{
		    c.JSON(400, gin.H{"error": err.Error()})
          return
	}
}