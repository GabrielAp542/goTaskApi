package dependencies

import (
	"github.com/GabrielAp542/goTaskApi/handler/http"
	"github.com/GabrielAp542/goTaskApi/internal/repositories"
	usecase "github.com/GabrielAp542/goTaskApi/internal/usecases"
	"gorm.io/gorm"
)

func DependenciesInjection(db *gorm.DB) *http.TaskHandler {
	taskRepository := repositories.NewTaskRepository(db)
	taskUseCase := usecase.NewTaskUseCase(taskRepository)
	taskHandler := http.NewTaskHandler(*taskUseCase)
	return taskHandler
}
