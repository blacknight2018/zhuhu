package orm

import (
	"demo/configure"
	"demo/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB

func init() {
	_db, err := gorm.Open("mysql", configure.GetDSN())
	if err != nil {
		logger.DBLog(logrus.Fields{}, logrus.ErrorLevel, "connection db error:"+err.Error())
	} else {
		_db.DB().SetMaxOpenConns(100)
		_db.DB().SetMaxIdleConns(20)
		_db.LogMode(true)
		db = _db
	}
}

func GetDB() *gorm.DB {
	return db
}
