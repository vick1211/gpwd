package main

import (
	"database/sql"
	"fmt"
	"gpwd/config"
	"gpwd/internal/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	conf := config.New()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password, conf.DB.DBName, conf.DB.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connected")

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	if err := database.Migrate(db); err != nil {
		log.Printf("Migration warning: %v", err)
	} else {
		log.Println("migrated succesfully")
	}
	
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	if err := r.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}