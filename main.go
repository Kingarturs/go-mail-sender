package main

import (
	"log"
	"fmt"
	"os"
	"net/http"
	"github.com/joho/godotenv"
	"./mail"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func IndexView() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
	return fn
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error cargando archivo .env")
	}
	PORT := os.Getenv("PORT")

	// Router
	router := gin.Default()
	router.Use(cors.Default())
	router.LoadHTMLGlob("./static/*")
	router.Use(static.Serve("/", static.LocalFile("./static", true)))

	// Routes
	router.GET("/", IndexView())
	router.POST("/send", mail.SendSimpleEmail())
	// router.Static("./static", "./static")

	fmt.Println("Go server is running in", PORT, "...")
	log.Fatal(router.Run(":"+PORT))
}
