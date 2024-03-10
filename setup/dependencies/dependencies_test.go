package dependencies

import (
	"testing"

	"gorm.io/gorm"
)

func TestDependencyInjection(t *testing.T) {
	handlerTest := DependenciesInjection(&gorm.DB{})
	if handlerTest == nil {
		t.Errorf("There is no task handler: %v", handlerTest)
	}
}
