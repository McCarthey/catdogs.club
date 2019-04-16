package libs

import (
	"net/smtp"
	"strings"
)

const (
	user   = "18836617@qq.com"
	passwd = "mjghzuuvuwmfcadg"
	host   = "smtp.qq.com"
	addr   = "smtp.qq.com:25"
)

func SendMail(to, sub, content string) error {
	auth := smtp.PlainAuth("", user, passwd, host)
	to_list := []string{to}
	nickname := "CatDogs"
	subject := sub
	content_type := "Content-Type: text/plain; charset=UTF-8"
	body := content
	msg := []byte("To: " + strings.Join(to_list, ",") + "\r\nFrom: " +
		nickname + "<" + user + ">\r\nSubject: " + subject + "\r\n" +
		content_type + "\r\n\r\n" + body)
	err := smtp.SendMail(addr, auth, user, to_list, msg)
	return err
}
