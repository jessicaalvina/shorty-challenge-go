package services

import (
	"os"
	"fmt"
	"net/url"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"testing"
	"github.com/joho/godotenv"
) 
var db *gorm.DB

func init() {

	// load env
	godotenv.Load("../.env")

	// db init
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")

	defaultTimezone := os.Getenv("SERVER_TIMEZONE")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=2&loc=%s",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
		url.QueryEscape(defaultTimezone),
	)

	db, err := gorm.Open("mysql", connection)
	if nil != err {
		panic(err)
	}

	db.LogMode(true)

	defer db.Close()

}

func TestV1ShortyService_GetByShortcodeSuccess(t *testing.T) {
	repository := V1ShortyServiceHandler(db.Begin())
	_, err := repository.GetByShortcode("google")
	
	if err != nil {
		t.Errorf("Success. Google shortcode found in the system")
	}
}

func TestV1ShortyService_GetByShortcodeFailed(t *testing.T) {
	repository := V1ShortyServiceHandler(db.Begin())
	_, err := repository.GetByShortcode("lollipop")
	
	if err == nil {
		t.Errorf("Failed. lollipop shortcode cant be found in the system")
	}
}

func TestV1ShortyService_GetByShortcodeStatsSuccess(t *testing.T) {
	repository := V1ShortyServiceHandler(db.Begin())
	_, err := repository.GetByShortcode("google")
	
	if err != nil {
		t.Errorf("Success. Showing google shortcode stats")
	}

}

func TestV1ShortyService_GetByShortcodeStatsFailed(t *testing.T) {
	repository := V1ShortyServiceHandler(db.Begin())
	_, err := repository.GetByShortcodeStats("lollipop")
	
	if err == nil {
		t.Errorf("Failed. lollipop shortcode cant be found in the system")
	}

}