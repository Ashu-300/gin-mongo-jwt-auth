package controllers

import (
	"context"
	"encoding/json"
	"gin-jwt-auth/database"
	"gin-jwt-auth/jwt"
	"gin-jwt-auth/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context){
	var data models.User
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		log.Println(err.Error())
	}
	collection := database.GetUserCollection("user")

	HashedPassword , err := bcrypt.GenerateFromPassword([]byte(data.Password) , bcrypt.DefaultCost) 
	if err != nil {
		log.Println(err.Error())
		return 
	}

	data.Password = string(HashedPassword)
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	result , err := collection.InsertOne(ctx , data)
	if err != nil {
		log.Println(err.Error())
	}
	if result == nil {
		c.JSON(500,gin.H{
			"message":"internal server error",
		})
		return
	}
	c.JSON(201,gin.H{
		"message":"user registered successfully",
		"user_detail": result ,
	})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}


func Login(c *gin.Context){
	var req LoginRequest

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest , gin.H{
			"message":"incomplete or wrong data",
		})
	}
	token := jwt.SignToken(req.Email)
	collection := database.GetUserCollection("user")
	var user models.User
	err = collection.FindOne(context.TODO(), map[string]interface{}{"email": req.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid email or password",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid email or password",
		})
		return
	}

	c.JSON(200,gin.H{
		"message":"user logged in successfully",
		"token":token,
	})
}
