package models

// AutoMigrations func automate database schema migrations
func AutoMigrations(db *Database) {
	db.Debug().AutoMigrate(&Sensor{})
}
