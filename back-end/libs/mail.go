package libs

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"

	configs "catdogs.club/back-end/configs/common"
)

var (
	email  = configs.FromEmail
	passwd = configs.EmailPasswd
)

const (
	host = "smtp.exmail.qq.com"
	addr = "smtp.exmail.qq.com:465"
	port = 465
)

func SendMail(to, sub, content string) error {
	header := make(map[string]string)
	header["From"] = "CatDogs" + "<" + email + ">"
	header["To"] = to
	header["Subject"] = sub
	header["Content-Type"] = "text/html; charset=UTF-8"
	body := content
	msg := ""
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + body
	auth := smtp.PlainAuth(
		fmt.Sprintf("%s:%d", host, port),
		email,
		passwd,
		host,
	)
	err := SendTLSMail(addr, auth, email, []string{to}, []byte(msg))
	return err
}

func Dial(add string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Println("Dialing TLS Failed: ", err)
		return nil, err
	}
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func SendTLSMail(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
	conn, err := Dial(addr)
	if err != nil {
		log.Println("Create smpt client failed: ", err)
		return err
	}
	defer conn.Close()
	if auth != nil {
		if ok, _ := conn.Extension("AUTH"); ok {
			if err = conn.Auth(auth); err != nil {
				log.Println("Error during AUTH", err)
				return err
			}
		}
	}

	if err = conn.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = conn.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := conn.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return conn.Quit()
}
