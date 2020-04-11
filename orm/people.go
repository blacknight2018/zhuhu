package orm

import (
	"demo/logger"
	"github.com/sirupsen/logrus"
)

type People struct {
	UserToken      string `gorm:"column:user_token;unique_index;PRIMARY_KEY"`
	Locations      string `gorm:"column:locations"`
	School         string `gorm:"column:school"`
	FollowerCount  int    `gorm:"column:follower_count"`
	FollowingCount int    `gorm:"column:following_count"`
	WorkIn         string `gorm:"column:workin"`
	Major          string `gorm:"column:major"`
}

func (u People) TableName() string {
	return "peoples"
}

func InsertPeople(u People) {
	if u.UserToken == "" {
		return
	}
	err := GetDB().Create(&u).Error
	if err != nil {
		logger.DBLog(logrus.Fields{}, logrus.InfoLevel, "insert error:"+err.Error())

	} else {
		logger.DBLog(logrus.Fields{}, logrus.InfoLevel, "insert success")
	}
}
func SelectUserSizeByConditional(conditional string) (RetSet []People) {
	GetDB().Model(&People{}).Where(conditional).Scan(&RetSet)
	return
}
