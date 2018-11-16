package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
	"ralali.com/controllers"
	"ralali.com/middleware"
)

func init() {

	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {

	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")

	defaultTimezone := os.Getenv("SERVER_TIMEZONE")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=1&loc=%s",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
		url.QueryEscape(defaultTimezone),
	)

	fmt.Println(connection)

	db, err := gorm.Open("mysql", connection)
	if nil != err {
		log.Fatal(err)
		os.Exit(1)
	}

	defer db.Close()

	defaultMiddleware := middleware.DefaultMiddleware{
		DB: db,
	}

	router := gin.Default()
	router.Use(defaultMiddleware.CORSMiddleware())

	controllers.V1UserControllerHandler(router, db)

	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")

	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	fmt.Println(serverString)

	router.Run(serverString)

}
