package dao

type Score struct {
	UserId int32 `gorm:"column:userId;comment:'用户id'"`
	Math float64 `gorm:"column:math;comment:'数学成绩'"`
	Chinese float64 `gorm:"column:chinese;comment:'中文成绩'"`
	English float64 `gorm:"column:english;comment:'英语成绩'"`
}

func (Score) TableName() string {
	return "score"
}
