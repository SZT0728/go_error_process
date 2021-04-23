package dao

type User struct {
	Id   int32  `gorm:"column:id;primary_key;comment:'用户id'"`
	Name string `gorm:"column:name;comment:'姓名'"`
	Age  int32  `gorm:"column:age;comment:'年龄'"`
}

func (User) TableName() string  {
	return "user"
}


