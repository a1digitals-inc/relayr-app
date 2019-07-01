package models

import (
	"log"

	"github.com/jinzhu/gorm"
	// This is necessary for gorm dialects
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Sensor struct
type Sensor struct {
	gorm.Model
	// ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);unique_index" json:"name"`
	Type string `gorm:"type:varchar(40);not null;unique_index" json:"type"`
	// CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	// UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

// NewSensor add new sensor to database
func (s *Sensor) NewSensor(db *Database) (interface{}, error) {
	rs := db.Debug().Create(&s)
	return rs.Value, rs.Error
}

// UpdateSensor update sensor name
func (s *Sensor) UpdateSensor(db *Database) (interface{}, error) {
	var erromsg = "Name or Type is empty!"
	if s.Name != "" {
		rs := db.Where("id = ?", s.ID).Find(&Sensor{}).Update("name", s.Name)
		return rs.Value, rs.Error
	}
	if s.Type != "" {
		rs := db.Where("id = ?", s.ID).Find(&Sensor{}).Update("type", s.Type)
		return rs.Value, rs.Error
	}
	return erromsg, log.Output(422, erromsg)
}

// GetAll sensors
func GetAll(db *Database) interface{} {
	return db.Order("id asc").Find(&[]Sensor{}).Value
}

// GetByID sensors
func GetByID(id string, db *Database) interface{} {
	if db.Where("id =?", id).First(&Sensor{}).RecordNotFound() {
		return nil
	}
	return db.Where("id = ?", id).First(&Sensor{}).Value
}

// Delete sensors from database
func Delete(id string, db *Database) (interface{}, error) {
	var rs *gorm.DB

	rs = db.Where("id = ?", id).Delete(&Sensor{})
	return rs.Value, rs.Error
}
