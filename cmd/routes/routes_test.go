package routes

import (
	"net/http"
	"testing"

	"github.com/GabrielAp542/goTask/cmd/dependencies"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateRoutes(t *testing.T) {
	handler := dependencies.DependenciesInjection(&gorm.DB{})
	CreateRoutes(handler)
	_, err := http.NewRequest("GET", "/tasks", nil)
	assert.NoError(t, err)
}
