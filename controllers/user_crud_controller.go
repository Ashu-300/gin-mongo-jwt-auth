package controllers

import (
	"context"
	"gin-jwt-auth/database"
	"gin-jwt-auth/jwt"
	"gin-jwt-auth/models"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)



func Profile(c *gin.Context) {
	collection := database.GetCollection("user")


	
	userInfo, exist := c.Get("userInfo")
	
	if !exist {
		c.JSON(500, gin.H{"error": "User not found in context"})
		return
	}

	// Type assert to *jwt.Claims
	claims, ok := userInfo.(*jwt.Claims)
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid user info type"})
		return
	} 

	email := claims.Email

	// Query MongoDB for this user
	var user models.User
	filter := bson.M{"email": email}

	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Println(err)
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Exclude password from response
	user.Password = ""

	c.JSON(200, gin.H{"user": user})
}
