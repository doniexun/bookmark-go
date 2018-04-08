package model

type User struct {
	Model // 自定义的model

	Mail		string	`gorm:"unique_index;not null;type:varchar(50)" json:"mail"`
	Password	string	`gorm:"not null;type:varchar(50)" json:"password"`
}

func (User) TableName() string { // 自定义表名
	return "user"
}

func AddUser(mail string, pwd string) bool {
	users := User{Mail: mail, Password: pwd}
	db.Create(&users)

	return true
}
