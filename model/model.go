package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Score struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserID   uint   `json:"userId" gorm:"column:user_id"`
	ExamName string `json:"exam" gorm:"column:exam_name"`
	Result   int    `json:"result"`
	Total    int    `json:"total"`
}
