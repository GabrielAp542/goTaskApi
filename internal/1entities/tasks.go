package entities

// Task struct, defines the table fields
type Task struct {
	TaskId     int    `gorm:"primaryKey" json:"task_id"`
	Task_name  string `json:"task_name"`
	Compleated bool   `json:"compleated"`
}
