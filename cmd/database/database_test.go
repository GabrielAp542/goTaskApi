package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnection(t *testing.T) {
	_, err := Conection("172.18.0.2",
		"postgres",
		"1234",
		"test_tasksDB",
		"5432")
	assert.NoError(t, err)
	/*if err != nil {
		t.Errorf("the database conection has failed, closing api. Error log: %v", err)
	}*/

	_, errf := Conection("uwu",
		"postgres",
		"1234",
		"test_tasksDB",
		"5432")
	assert.Error(t, errf)
	/*if errf == nil {
		t.Error("The conection was succesfull when It wasn't expected")
	}*/
}
