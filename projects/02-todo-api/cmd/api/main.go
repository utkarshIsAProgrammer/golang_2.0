package main

import (
	"log"
	"todo-api/internal/config"
	"todo-api/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	var cfg *config.Config
	var err error
	cfg, err = config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
	}

	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabseURL)
	if err != nil {
		log.Fatal("Failed to connect to the PostgreSQL database!", err)
	}
	defer pool.Close()

	// gin instance
	var router *gin.Engine = gin.Default()
	router.SetTrustedProxies(nil) // no proxies allowed

	// home route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "TODO API is running!",
			"status":   "success",
			"database": "connected",
		})
	})

	router.Run(":" + cfg.Port)
}
