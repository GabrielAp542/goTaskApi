package dependencies

import (
	"github.com/GabrielAp542/goTask/handler/http"
	"github.com/GabrielAp542/goTask/internal/repositories"
	usecase "github.com/GabrielAp542/goTask/internal/usecases"
	"gorm.io/gorm"
)

func DependenciesInjection(db *gorm.DB) *http.TaskHandler {
	taskRepository := repositories.NewTaskRepository(db)
	taskUseCase := usecase.NewTaskUseCase(taskRepository)
	taskHandler := http.NewTaskHandler(*taskUseCase)
	return taskHandler
}
