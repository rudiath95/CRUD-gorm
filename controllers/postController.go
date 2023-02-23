package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rudiath95/CRUD-gorm/initializers"
	"github.com/rudiath95/CRUD-gorm/models"
)

func PostCreate(c *gin.Context) {

	//Call DB
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	//Create Query
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to post",
		})
		return
	}

	//Return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostGet(c *gin.Context) {

	//Call DB
	var post []models.Post
	initializers.DB.Find(&post)

	//Return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostFind(c *gin.Context) {
	//get id
	id := c.Param("id")

	//Call DB
	var post models.Post
	initializers.DB.Find(&post, id)

	//Return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	//get id
	id := c.Param("id")

	//Call DB
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)
	var post models.Post
	initializers.DB.Find(&post, id)

	//Update
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//Return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	//get id
	id := c.Param("id")

	//Delete
	initializers.DB.Delete(&models.Post{}, id)

	//Return
	c.JSON(200, gin.H{
		"post": "Deleted",
	})
}
