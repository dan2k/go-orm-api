package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"example/go-orm-api/model"
	"github.com/gin-contrib/cors"
)

func main() {
	dsn := "mitsu:secret@tcp(db:3306)/go_local?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		var users []model.User 
		db.Find(&users)
		c.JSON(http.StatusOK, users)
	})
	r.GET("/users/:id", func(c *gin.Context) {
		id :=c.Param("id")
		var user model.User 
		db.First(&user,id)
		c.JSON(http.StatusOK, user)
	})
	r.POST("/users",func(c *gin.Context){
		var user model.User
		if err := c.ShouldBindJSON(&user); err !=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		result :=db.Create(&user)
		c.JSON(200,gin.H{"RowAffected":result.RowsAffected})
	})
	r.DELETE("/users/:id",func(c *gin.Context){
		id :=c.Param("id")
		user :=model.User{}
		db.First(&user,id)
		db.Delete(&user)
		c.JSON(http.StatusOK,user)	
	})
	r.PUT("/users/:id",func(c *gin.Context){
		user := model.User{}
		if err :=c.ShouldBindJSON(&user);err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		updatedUser :=model.User{}
		db.First(&updatedUser,user.ID)
		updatedUser.Fname=user.Fname
		updatedUser.Lname=user.Lname
		updatedUser.Username=user.Username
		updatedUser.Avatar=user.Avatar
		db.Save(&updatedUser)
		c.JSON(200,updatedUser)

	})
	r.Use(cors.Default())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
