package Models

import Mysql "gin/Databases"

type Email struct {
	ID         int
	UserID     int      `gorm:"index"`                            // 外键 (属于), tag `index`是为该列创建索引
	Email      string   `gorm:"type:varchar(100);unique_index"`   // `type`设置sql类型, `unique_index` 为该列设置唯一索引
	Subscribed bool
}


func (this *Email) Insert() (id int, err error) {
	result := Mysql.DB.Create(&this)
	id = this.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}


func (this *Email) Delete() (id int, err error) {
	result := Mysql.DB.Delete(&this)
	id = this.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}


func (this *Email) Update() (id int, err error) {
	result := Mysql.DB.Save(&this)
	id = this.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
