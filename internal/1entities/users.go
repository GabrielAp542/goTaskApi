package entities

// Task struct, defines the table fields
type Users struct {
	UserId   int    `gorm:"primaryKey" json:"user_id"`
	Username string `json:"username"`
}
