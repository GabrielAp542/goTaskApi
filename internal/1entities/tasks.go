package entities

// Task struct, defines the table fields
type Task struct {
	TaskId    int    `gorm:"primaryKey" json:"task_id"`
	Task_name string `json:"task_name"`
	Completed bool   `json:"Completed"`
	Id_User   *int   `json:"user_id"`
	User      Users  `gorm:"foreignKey:user_id"`
}
