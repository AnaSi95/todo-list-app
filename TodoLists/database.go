package TodoLists

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectToDB() *gorm.DB {

	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=todo_db password=qwerty sslmode=disable")
	if err != nil {
		panic("Connection failed!")
	}
	db.AutoMigrate(&Item{})

	return db
}
