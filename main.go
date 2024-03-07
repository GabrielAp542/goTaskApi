// main.go
package main

import (
	"log"

	"github.com/GabrielAp542/goTask-Api-Gabriel/cmd/database"
	"github.com/GabrielAp542/goTask-Api-Gabriel/cmd/dependencies"
	"github.com/GabrielAp542/goTask-Api-Gabriel/cmd/routes"
	_ "github.com/GabrielAp542/goTask-Api-Gabriel/docs"
)

// @title Task Api Go
// @version 1.0
// @description Task Api
// @host localhost:8080
// @basepath /
func main() {
	//database connection
	/*db, err := database.Conection(os.Getenv("DB_HOST"),
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),
	os.Getenv("DB_PORT"))*/
	db, err := database.Conection("10.0.1.3",
		"postgres",
		"1234",
		"task_apiDB",
		"5432")
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
