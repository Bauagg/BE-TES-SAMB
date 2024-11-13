package main

import (
	"TES-SAMB-GO/config"
	"TES-SAMB-GO/databases"
	"TES-SAMB-GO/midleware"
	"TES-SAMB-GO/migration"
	"TES-SAMB-GO/routes"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// config
	config.ConfigEnv()

	// connect databases
	databases.Connect()
	migration.Migration()

	app := gin.Default()

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodPatch,
			http.MethodHead,
			http.MethodConnect,
			http.MethodTrace,
		},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	// midleware
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(midleware.ErrorMiddleware())

	// Router
	routes.IndexRouter(app)

	err := app.Run(config.APP_PORT)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
