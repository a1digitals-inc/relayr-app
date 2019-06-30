package models

import "github.com/jinzhu/gorm"

// Define which struct to perform operations
const (
	USERS     = "users"
	FEEDBACKS = "feedbacks"
)

// GetAll users for feedbacks
func GetAll(table string) interface{} {
	db := Connect()
	defer db.Close()
	switch table {
	case USERS:
		return db.Order("id asc").Find(&[]User{}).Value
	case FEEDBACKS:
		//http://doc.gorm.io/associations.html#has-many
		var feedbacks []Feedback
		db.Order("id asc").Find(&feedbacks)
		for i := range feedbacks {
			db.Model(feedbacks[i]).Related(&feedbacks[i].User)
		}
		return feedbacks
	}
	return nil
}

// GetByID users or feedback
func GetByID(table, id string) interface{} {
	db := Connect()
	defer db.Close()
	switch table {
	case USERS:
		if db.Where("id =?", id).First(&User{}).RecordNotFound() {
			return nil
		}
		return db.Where("id = ?", id).First(&User{}).Value
	case FEEDBACKS:
		return db.Where("id = ?", id).First(&Feedback{}).Value
	}
	return nil
}

// Delete users or feedback from database
func Delete(table, id string) (int64, error) {
	db := Connect()
	defer db.Close()
	var rs *gorm.DB
	switch table {
	case USERS:
		rs = db.Where("id = ?", id).Delete(&User{})
		break
	case FEEDBACKS:
		rs = db.Where("id = ?", id).Delete(&Feedback{})
		break
	default:
		return 0, nil
	}
	return rs.RowsAffected, rs.Error
}
