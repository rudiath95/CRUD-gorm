package initializers

import "github.com/rudiath95/CRUD-gorm/models"

func MigrateDatabases() {
	DB.AutoMigrate(models.Post{})
}
