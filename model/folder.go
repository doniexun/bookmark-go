package model

type Folders struct {
	Model

	Name		string	`gorm:"not null;type:varchar(255)";DEFAULT:'Unnamed' json:"name"`
	UserId		int		`gorm:"not null;" json:"user_id"`
}

func (Folders) TableName() string { // 自定义表名
  	return "folder"
}

func AddFolder(name string, userId int) bool {
	db.Create(&Folders{
		Name: name,
		UserId: userId,
	})

	return true
}

// func GetFolders(pageNum int, pageSize int, maps interface {}) (folders []Folder) {
// 	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&folders)

// 	return
// }

func GetFolderTotal(maps interface {}) int {
	var count int
	db.Model(&Folders{}).Where(maps).Count(&count)
	return count
}
