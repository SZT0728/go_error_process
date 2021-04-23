package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var db *gorm.DB

func Init(dbPath string) error {
	var err error
	db, err = gorm.Open("mysql", dbPath)
	if err != nil {
		return errors.Wrapf(err, "数据库链接为:%s", dbPath)
	}
	err = db.AutoMigrate(&User{}).Error
	if err != nil {
		return errors.Wrap(err, "user AutoMigrate failed")
	}

	err = db.AutoMigrate(&Score{}).Error
	if err != nil{
		return errors.Wrap(err,"score AutoMigrate failed")
	}

	return nil
}

func GetUserById(uid int32) (*User, error) {
	var user User
	err := db.Model(&User{}).Where("id=?", uid).Find(&user).Error
	if err != nil {
		return nil, errors.Wrapf(err, "获取user失败:userId=%d",uid)
	}
	return &user, nil
}

func GetScoreByUserId(uid int32) (*Score, error) {
	var score Score
	err := db.Model(&Score{}).Where("userId=?").Find(&score).Error
	if err != nil{
		return nil, errors.Wrapf(err, "获取score失败:userId=%d",uid)
	}
	return &score, nil
}


func AddUserAndScore(userName string, age int32, math, chinese, english float64) (int32, error) {
	user := &User{
		Name: userName,
		Age:  age,
	}

	tx := db.Begin()

	err := tx.Model(&User{}).Create(user).Error
	if err != nil {
		tx.Rollback()
		return 0, errors.Wrap(err, "创建用户失败")
	}

	score := &Score{
		UserId:user.Id,
		Math:math,
		Chinese:chinese,
		English:english,
	}

	err = tx.Model(&Score{}).Create(&score).Error
	if err != nil{
		tx.Rollback()
		return 0, errors.Wrap(err, "增加分数失败")
	}
	tx.Commit()
	return user.Id, nil
}
