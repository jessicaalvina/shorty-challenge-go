package repositories

import (
	"os"
	"fmt"
	"net/url"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"testing"
	"github.com/joho/godotenv"
	"../objects"
	// "github.com/jinzhu/gorm"
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

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=1&loc=%s",
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

func TestV1ShortyRepository_GetByShortcodeFailed(t *testing.T) {

	repository := V1ShortyRepositoryHandler(db.Begin())
	_, err := repository.GetByShortcode("example")
	
	if err == nil {
		t.Errorf("Failed. Example shortcode cant be found in the system")
	}
}

func TestV1ShortyRepository_GetByShortcodeSuccess(t *testing.T) {

	repository := V1ShortyRepositoryHandler(db.Begin())
	_, err := repository.GetByShortcode("google")
	
	if err != nil {
		t.Errorf("Success. Google shortcode found in the system, redirecting...")
	}
}

func TestV1ShortyRepository_UpdateRedirectCountFailed(t *testing.T) {
	repository := V1ShortyRepositoryHandler(db.Begin())
	postObject := objects.V1ShortyObjectRequest{Shortcode: "bingobingo"}

	errBool := repository.UpdateRedirectCount("apple", postObject)
	
	if errBool == false {
		t.Errorf("Failed. Can't update redirect count for apple shortcode")
	}
}

func TestV1ShortyRepository_UpdateRedirectCountSuccess(t *testing.T) {
	repository := V1ShortyRepositoryHandler(db.Begin())
	postObject := objects.V1ShortyObjectRequest{Shortcode: "google"}

	errBool := repository.UpdateRedirectCount("google",postObject)
	
	if errBool != false {
		t.Errorf("Success. Add 1 for redirect count for google shortcode")
	}
}

func TestV1ShortyRepository_GetByShortcodeStatsFailed(t *testing.T) {

	repository := V1ShortyRepositoryHandler(db.Begin())
	_, err := repository.GetByShortcodeStats("example")
	
	if err == nil {
		t.Errorf("Failed. Example shortcode cant be found in the system")
	}
}

func TestV1ShortyRepository_GetByShortcodeStatsSuccess(t *testing.T) {

	repository := V1ShortyRepositoryHandler(db.Begin())
	_, err := repository.GetByShortcodeStats("google")
	
	if err != nil {
		t.Errorf("Success. Show google shortcode stats")
	}
}