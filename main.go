package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rudiath95/CRUD-gorm/controllers"
	"github.com/rudiath95/CRUD-gorm/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnecttoDB()
	initializers.MigrateDatabases()
}

func main() {
	r := gin.Default()
	r.POST("/post", controllers.PostCreate)
	r.GET("/post", controllers.PostGet)
	r.GET("/post/:id", controllers.PostFind)
	r.PUT("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)

	r.Run()
}
