package database

import (
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	_, err := Conection("172.18.0.2",
		"postgres",
		"1234",
		"test_tasksDB",
		"5432")
	if err != nil {
		t.Errorf("the database conection has failed, closing api. Error log: %v", err)
	}

	_, errf := Conection("uwu",
		"postgres",
		"1234",
		"test_tasksDB",
		"5432")
	if errf == nil {
		t.Error("The conection was succesfully when I wasn't expected")
	}
}
