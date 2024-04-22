package platform

import (
	"net/http"

	"i9sh/platform/login"
	"i9sh/platform/middleware"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

func New(auth *auth.Client) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	router.Static("/static", "./static")

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {

		uid := "None"

		tokenID, err := login.ExtractUIDFromSession(c, auth)
		if err == nil {
			uid = tokenID
		}

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Title":  "Public Page",
			"Header": "Welcome to Public Page",
			"UID":    uid,
			"Year":   2024,
		})
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", nil)
	})

	router.POST("/verifyToken", login.VerifyToken(auth))

	return router
}
