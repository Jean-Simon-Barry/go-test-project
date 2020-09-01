package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
	"github.com/spf13/viper"

	"unity-go-project/book"
)

// Init : Initializes the db
func Init() *gorm.DB {
	viper.AutomaticEnv()
	viperUser := viper.Get("POSTGRES_USER")
	viperPassword := viper.Get("POSTGRES_PASSWORD")
	viperDb := viper.Get("POSTGRES_DB")
	viperHost := viper.Get("POSTGRES_HOST")
	viperPort := viper.Get("POSTGRES_PORT")
	postgresConn := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viperHost, viperPort, viperUser, viperDb, viperPassword)
	fmt.Println("conname is\t\t", postgresConn)
	db, err := gorm.Open("postgres", postgresConn)
	if err != nil {
		log.Println(err)
		panic("Failed to connect to database! ")
	}
	db.AutoMigrate(&book.Book{})
	m := book.Book{Author: "JSB!!!", Title: "Heard the nerd herd"}
	log.Println(m)
	db.Create(&m)
	db.Create(&book.Book{Author: "JSBwatt!!!", Title: "Heard the nerd herd"})

	return db
}
