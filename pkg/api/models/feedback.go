package models

import "time"

// Feedback struct
type Feedback struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Comment   string    `gorm:"type:varchar(255)" json:"comment"`
	UserID    uint32    `json:"user_id"`
	User      User      `json:"user"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

// NewFeedback store on database
func NewFeedback(f Feedback) (int64, error) {
	db := Connect()
	defer db.Close()
	rs := db.Create(&f)
	return rs.RowsAffected, rs.Error
}
