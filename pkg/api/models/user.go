package models

import "time"

// User struct
type User struct {
	ID        uint32     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string     `gorm:"type:varchar(20);unique_index" json:"name"`
	Email     string     `gorm:"type:varchar(40);not null;unique_index" json:"email"`
	CreatedAt time.Time  `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:current_timestamp()" json:"updated_at"`
	Feedbacks []Feedback `gorm:"ForeignKey:UserId" json:"feedbacks"`
}

// NewUser add new user to datavase
func NewUser(user User) (int64, error) {
	db := Connect()
	defer db.Close()
	rs := db.Create(&user)
	return rs.RowsAffected, rs.Error
}

// UpdateUser update user info
func UpdateUser(user User) (int64, error) {
	db := Connect()
	defer db.Close()
	rs := db.Where("id = ?", user.ID).Find(&User{}).Update("name", user.Name)
	return rs.RowsAffected, rs.Error
}
