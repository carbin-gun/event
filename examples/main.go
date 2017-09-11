package main

import (
	"fmt"

	"github.com/carbin-gun/event"
)

func main() {
	exampleA()
	fmt.Println("--------------")
	exampleB()
}

func exampleA() {
	e := event.New()
	e.AddListener("mail.push", func(mailto string, title string, content string) {
		fmt.Printf("mail to :%s,title:%s,content:%s\n", mailto, title, content)
	})
	e.AddListener("mail.push", func(mailto string, title string, content string) {
		fmt.Printf("after mail to :%s,title:%s,content:%s\n", mailto, title, content)
	})
	e.Fire("mail.push", "zhengchao.deng@meican.com", "hello,mail test", "hello,don't reply to this email,it's just a test")
}

//Mail definition for mail info
type Mail struct {
	MailTo  string
	Title   string
	Content string
	Sender  string
}

func exampleB() {
	e := event.New()
	e.AddListener("mail.push", func(mail Mail) {
		fmt.Printf("sending mail to :%s,title:%s,content:%s,sender:%s\n", mail.MailTo, mail.Title, mail.Content, mail.Sender)
	})
	e.AddListener("mail.push", func(mail Mail) {
		fmt.Printf("after mail to :%s,title:%s,content:%s,sender:%s\n", mail.MailTo, mail.Title, mail.Content, mail.Sender)
	})
	mail := Mail{"carbin-gun@github.com", "inquiry email", "do't reply", "no-reply@github.com"}
	e.Fire("mail.push", mail)
}
