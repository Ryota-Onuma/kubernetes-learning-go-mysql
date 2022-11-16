package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var allUsers []User

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	db, err := initDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var count int
	if err := db.Table("users").Count(&count).Error; err != nil {
		panic(err.Error())
	}

	newUser := newUser(count+1, fmt.Sprintf("ryota%d", count+1))
	if err := db.Create(&newUser).Error; err != nil {
		panic(err.Error())
	}

	if err := db.Find(&allUsers).Error; err != nil {
		panic(err.Error())
	}

	e.GET("/", top)
	e.Logger.Fatal(e.Start(":8080"))
}

func top(c echo.Context) error {
	return c.JSON(http.StatusOK, allUsers)
}

func initDB() (*gorm.DB, error) {
	dbms := "mysql"
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := fmt.Sprintf("tcp(%s:3306)", os.Getenv("MYSQL_HOST"))
	dbName := os.Getenv("MYSQL_DATABASE")

	connect := user + ":" + password + "@" + host + "/" + dbName
	// DBに接続
	db, err := gorm.Open(dbms, connect)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return db, nil
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func newUser(id int, name string) *User {
	return &User{
		ID:   id,
		Name: name,
	}
}
