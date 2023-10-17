package configs

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"tes_project_siesta/init/response/schema"
	"tes_project_siesta/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {

	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true&charset=utf8mb4", DB_USER, DB_PASS, DB_HOST, DB_PORT)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + DB_NAME)
	if err != nil {
		log.Fatal("Gagal membuat database: ", err)
	}

	_, err = db.Exec(fmt.Sprintf("USE %s", DB_NAME))
	if err != nil {
		log.Fatal("Gagal menggunakan database: ", err)
	}

	err = MigrateDB()
	if err != nil {
		log.Fatal("Gagal migrate table: ", err)
	}

}

func MigrateDB() (err error) {

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: DB,
	}), &gorm.Config{})

	if err != nil {
		return
	}

	migrate(gormDB)

	seedUser(gormDB)

	return
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&schema.User{},
		&schema.Loan{},
	)
}

func seedUser(db *gorm.DB) (err error) {

	var user schema.User

	err = db.First(&user, "id", 1).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {

		password, err := utils.HashPassword("Ahmad1")

		if err != nil {
			return err
		}

		user.Username = "Ahmad"
		user.Password = string(password)

		err = db.Create(&user).Error
		return err
	}

	return
}
