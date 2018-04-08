package model

type Bookmark struct {
	Model

	Title		string	`gorm:"not null;type:varchar(255)";DEFAULT:'Untitle' json:"title"`
	Url			string	`gorm:"not null;type:varchar(255)" json:"url"`
	UserId		int		`gorm:"not null;" json:"user_id"`
	FolderId	int		`gorm:"not null;" json:"folder_id"`
}

func (Bookmark) TableName() string { // 自定义表名
	return "bookmark"
}
