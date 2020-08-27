package Models

import (
	"gin/Databases"
)


type Email struct {
	ID      int		`json:"id"`
	UserID  int     `json:"user_id"     gorm:"index"`                            // 外键 (属于), tag `index`是为该列创建索引
	Email   string  `json:"email"       gorm:"type:varchar(100);unique_index"`   // `type`设置sql类型, `unique_index` 为该列设置唯一索引
	Subscribed bool `json:"subscribed"`
}


type Test struct {
	Id int
	Testcol string `gorm:"column:testcol"`
}

// 设置Test的表名为`test`
// func (Test) TableName() string {
//     return "test"
// }

func (this *Test) Insert() (id int, err error) {
	result := Mysql.DB.Create(&this)
	id = this.Id
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
