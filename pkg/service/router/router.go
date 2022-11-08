package router

import (
	"app/pkg/service/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())

	itemgroup := r.Group("v1/item")
	{
		itemgroup.GET("find/:name", controllers.FindItemsByName)
		itemgroup.GET("find/category/:category", controllers.FindItemsByCategory)
	}

	categorygroup := r.Group("v1/category")
	{
		categorygroup.GET("find/:name", controllers.FindCategoriesByName)
	}

	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
