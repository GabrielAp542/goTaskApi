// main.go
package main

import (
	"log"
	"os"

	"github.com/GabrielAp542/goTask/cmd/database"
	"github.com/GabrielAp542/goTask/cmd/dependencies"
	"github.com/GabrielAp542/goTask/cmd/routes"
	_ "github.com/GabrielAp542/goTask/docs"
	// gin-swagger middleware
	// swagger embed files
	// swagger embed files
)

// @title Task Api Go
// @version 1.0
// @description Task Api
// @host localhost:8080
// @basepath /
func main() {
	//database connection
	db, err := database.Conection(os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	if err != nil {
		log.Panicf("the database conection has failed, closing api. Error log: %v", err)
	}
	// Dependency Injection
	taskHandler := dependencies.DependenciesInjection(db)
	// configuración de rutas
	router := routes.CreateRoutes(taskHandler)
	// Iniciar servidor
	router.Run(":8080")

}
