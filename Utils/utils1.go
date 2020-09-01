package Utils

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)


const(
	ADDR = "smtp.163.com:25"
	IDENTITY = ""
	USERNAME = "sunckdebug@163.com"
	PASSWORD = "GLULZSPJZHAFQUKN"
	HOST = "smtp.163.com"
)


func SendMail(fromUser, toUser, subject string, filename string) error {
	// NewEmail返回一个email结构体的指针
	e := email.NewEmail()
	// 发件人
	e.From = fromUser
	// 收件人(可以有多个)
	e.To = []string{toUser}
	// 邮件主题
	e.Subject = subject
	// 以路径将文件作为附件添加到邮件中
	e.AttachFile(filename)
	// 发送邮件(如果使用QQ邮箱发送邮件的话，passwd不是邮箱密码而是授权码)
	return e.Send(ADDR, smtp.PlainAuth(IDENTITY, USERNAME, PASSWORD, HOST))
}