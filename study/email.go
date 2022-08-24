package main

import (
	"gopkg.in/gomail.v2"
)

func main() {
	smtpAddr := "smtp.163.com"
	smtpPort := 465
	smtpUser := "wxr_624@163.com"
	smtpPassword := "FZTPKITVUUZGGOWB"

	from := "wxr_624@163.com"
	tos := []string{"834555340@qq.com"}
	subject := "email of go program "
	content := "成功发送邮件"

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", tos...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)
	d := gomail.NewDialer(smtpAddr, smtpPort, smtpUser, smtpPassword)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	//m := gomail.NewMessage()
	//m.SetHeader("From", " alex@example . com")
	//m.SetHeader("To", "bob@example . com", "cora@example .com")
	//m.SetAddressHeader("CC", "dan@example. com", "Dan")
	//m.SetHeader("Subject", "Hello!")
	//m.SetBody("text/htm1", "Hello : <b>Bob</b>，and，<i>Cora</i>!")
	//m.Attach(" /home/Alex/lolcat.jpg")
	//d := gomail.NewDialer(" smtp. example.com", 587, "user", "123456 ")
	//
	//d := gomail.NewDialer("smtp.example.com", 587, "user", "123456")
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send emails using d.
}
