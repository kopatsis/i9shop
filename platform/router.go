package platform

import (
	"net/http"

	"i9sh/platform/middleware"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	router.LoadHTMLGlob("templates/*")
	router.GET("/public", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Title":   "Public Page",
			"Header":  "Welcome to Public Page",
			"Message": "Blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah",
			"Year":    2024,
		})
	})

	router.GET("/unprotected", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Title":   "Private Page",
			"Header":  "Welcome to Private Page",
			"Message": "Blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah",
			"Year":    2024,
		})
	})

	return router
}
