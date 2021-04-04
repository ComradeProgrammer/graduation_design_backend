package db

import (
	"graduation_design/internal/pkg/logs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func DBInit() {
	var err error
	db, err = gorm.Open("sqlite3", "app.db")
	if err != nil {
		logs.Error("Open database app.db failed")
		return
	}
	db.AutoMigrate(&ProjectDB{}, &Regex{})
	logs.Info("Open database app.db succeeded")
}
