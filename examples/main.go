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
	e := event.New("mail.push")
	e.AddListener(func(mailto string, title string, content string) {
		fmt.Printf("mail to :%s,title:%s,content:%s\n", mailto, title, content)
	})
	e.Fire("zhengchao.deng@meican.com", "test email", "don't reply")
}

//Mail definition for mail info
type Mail struct {
	MailTo  string
	Title   string
	Content string
	Sender  string
}

func exampleB() {
	e := event.New("mail.push")
	e.AddListener(func(mail Mail) {
		fmt.Printf("sending mail to :%s,title:%s,content:%s,sender:%s\n", mail.MailTo, mail.Title, mail.Content, mail.Sender)
	})
	e.AddListener(func(mail Mail) {
		fmt.Printf("after mail to :%s,title:%s,content:%s,sender:%s\n", mail.MailTo, mail.Title, mail.Content, mail.Sender)
	})
	mail := Mail{"carbin-gun@github.com", "inquiry email", "do't reply", "no-reply@github.com"}
	e.Fire(mail)
}
