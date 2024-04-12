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
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	return router
}
