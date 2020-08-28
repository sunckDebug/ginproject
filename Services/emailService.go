package Services


import "gin/Models"


type Email struct {
	Id         int     `json:"id"`
	UserID     int     `json:"user_id"`
	Email      string  `json:"email"`
	Subscribed bool    `json:"subscribed"`
}


func (this *Email) Insert() (id int, err error) {
	var emailModel Models.Email
	emailModel.ID = this.Id
	emailModel.UserID = this.UserID
	emailModel.Email = this.Email
	emailModel.Subscribed = this.Subscribed
	id, err = emailModel.Insert()
	return
}


func (this *Email) Delete() (id int, err error) {
	var emailModel Models.Email
	emailModel.ID = this.Id
	emailModel.UserID = this.UserID
	emailModel.Email = this.Email
	emailModel.Subscribed = this.Subscribed
	id, err = emailModel.Delete()
	return
}


func (this *Email) Update() (id int, err error) {
	var emailModel Models.Email
	emailModel.ID = this.Id
	emailModel.UserID = this.UserID
	emailModel.Email = this.Email
	emailModel.Subscribed = this.Subscribed
	id, err = emailModel.Update()
	return
}