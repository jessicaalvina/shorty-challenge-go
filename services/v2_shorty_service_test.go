package services

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	"../objects"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// var db *gorm.DB

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

func TestV2ShortyService_ValidateShortcodeSuccess(t *testing.T) {
	service := V2ShortyServiceHandler(db.Begin())
	errBool := service.ValidateShortcode("google")

	if errBool != false {
		t.Errorf("Success. google is valid shortcode")
	}
}

func TestV2ShortyService_ValidateShortcodeFailed(t *testing.T) {
	service := V2ShortyServiceHandler(db.Begin())
	errBool := service.ValidateShortcode("lollipop")

	if errBool == false {
		t.Errorf("Failed. lollipop is not valid shortcode")
	}
}

func TestV2ShortyService_ValidateUrlSuccess(t *testing.T) {
	service := V2ShortyServiceHandler(db.Begin())
	errBool := service.ValidateUrl("http://www.google.com")

	if errBool != false {
		t.Errorf("Success. http://www.google.com is valid Url")
	}
}

func TestV2ShortyService_ValidateUrlFailed(t *testing.T) {
	service := V2ShortyServiceHandler(db.Begin())
	errBool := service.ValidateUrl("ww.lollipopcom")

	if errBool == false {
		t.Errorf("Failed. ww.lollipopcom is not valid Url")
	}
}

func TestV2ShortyService_PostByShortenSuccess(t *testing.T) {
	service := V2ShortyServiceHandler(db.Begin())
	postObject := objects.V2ShortyObjectRequest{Shortcode: "yahhoo", Url: "http://www.yahoo.com"}

	_, err := service.PostByShorten(postObject)

	if err != nil {
		t.Errorf("Success. yahhoo shortcode created with http://www.yahoo.com url")
	}
}

func TestV2ShortyService_PostByShortenFailed(t *testing.T) {
	service := V2ShortyServiceHandler(db.Begin())
	postObject := objects.V2ShortyObjectRequest{Shortcode: "google", Url: "http://www.google.com"}

	_, err := service.PostByShorten(postObject)

	if err == nil {
		t.Errorf("Failed. google shortcode has been created before")
	}
}
