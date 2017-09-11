# event
event-driven design in go 


```go
func exampleA() {
	e := event.New()
	e.AddListener("mail.push", func(mailto string, title string, content string) {
		fmt.Printf("mail to :%s,title:%s,content:%s\n", mailto, title, content)
	})
	e.Fire("mail.push", "zhengchao.deng@meican.com", "test email", "don't reply")
}
```
Things can be done easily with listener-way. 
