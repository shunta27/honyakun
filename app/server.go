package app

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shunta27/honyakun/line"
)

// Start function
func Start() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := setup()
	router.Run(":" + port)
}

func setup() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.POST("/lineBotCallback", line.BotCallback)

	return router
}
