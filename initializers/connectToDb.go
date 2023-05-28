package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectToDb() {
	// postgres://gfjkdhvc:nlC5Xqr7NsxrRDcrjOxcPlydskbnNsWN@rajje.db.elephantsql.com/gfjkdhvc
	dsn := os.Getenv("DB")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db!")
	}
}
