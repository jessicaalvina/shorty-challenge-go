package repositories

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	"../objects"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	// "github.com/jinzhu/gorm"
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

func TestV2ShortyRepository_PostByShortenRegexFailed(t *testing.T) {
	repository := V2ShortyRepositoryHandler(db.Begin())
	postObject := objects.V2ShortyObjectRequest{Shortcode: "bingobingo", Url: "http://www.bingo.com"}

	_, err := repository.PostByShorten(postObject)

	if err == nil {
		t.Errorf("Failed. Shortcode is not meet following regexep ^[0-9a-zA-Z_]{4,}$")
	}
}

func TestV2ShortyRepository_PostByShortenUrlFailed(t *testing.T) {
	repository := V2ShortyRepositoryHandler(db.Begin())
	postObject := objects.V2ShortyObjectRequest{Shortcode: "bottle", Url: "w.bottlecom"}

	_, err := repository.PostByShorten(postObject)

	if err == nil {
		t.Errorf("Failed. Url is not present")
	}
}

func TestV2ShortyRepository_PostByShortenDuplicateFailed(t *testing.T) {
	repository := V2ShortyRepositoryHandler(db.Begin())
	postObject := objects.V2ShortyObjectRequest{Shortcode: "google", Url: "http://www.google.com"}

	_, err := repository.PostByShorten(postObject)

	if err == nil {
		t.Errorf("Failed. google shortcode has been created before")
	}
}

func TestV2ShortyRepository_PostByShortenSuccess(t *testing.T) {
	repository := V2ShortyRepositoryHandler(db.Begin())
	postObject := objects.V2ShortyObjectRequest{Shortcode: "yahhoo", Url: "http://www.yahoo.com"}

	_, err := repository.PostByShorten(postObject)

	if err != nil {
		t.Errorf("Success. yahhoo shortcode created with http://www.yahoo.com url")
	}
}
