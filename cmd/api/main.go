// /cmd/api/main.go

package main

import (
	"ginTemplate/pkg/config"
	"ginTemplate/pkg/logger"
	"ginTemplate/pkg/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	// Load environment variables
	err := config.LoadConfig("./pkg/config")
	if err != nil {
		logger.Fatal("Could not load config", err)
	}

	router := gin.New()

	// Attach middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Initialize routes
	routes.InitializeRoutes(router)

	// Set trusted proxies (modify as needed)
	err = router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		logger.Fatal("Could not set trusted proxies", err)
	}

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = viper.GetString("PORT")
	}
	router.Run(":" + port)
}
