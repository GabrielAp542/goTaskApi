package database

import (
	"fmt"

	"github.com/GabrielAp542/goTask/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Conection(host string, user string, password string, dbname string, port string) (*gorm.DB, error) {
	// Postgres conection by getting env variables
	dsn := fmt.Sprintf("host=%s user=%s  password=%s  dbname=%s  port=%s",
		host, user, password, dbname, port)
	//open conection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//detect any error
	if err != nil {
		return db, err
	}
	// Migrar esquemas
	db.AutoMigrate(&entities.Task{})
	db.AutoMigrate(&entities.Users{})
	return db, err
}

// creats connection with testing database
func TestingDB(fail bool) (*gorm.DB, error) {
	var host string
	if fail {
		host = "invalid"
	} else {
		host = "172.24.0.3"
	}
	db, err := Conection(host,
		"postgres",
		"1234",
		"test_tasksDB",
		"5432")
	return db, err
}
