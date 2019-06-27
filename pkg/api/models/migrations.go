package models

// AutoMigrations func automate database schema migrations
func AutoMigrations() {
	db := Connect()
	defer db.Close()
	db.DropTableIfExists(&Feedback{}, &User{})
	db.Debug().AutoMigrate(&User{}, &Feedback{})
	db.Model(&Feedback{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
}
