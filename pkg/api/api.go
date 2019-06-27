package api

import (
	"github.com/andrleite/relayr-app/pkg/api/models"
)

// Run database migrations and call server listen
func Run() {
	models.AutoMigrations()
	listen(9000)
}
