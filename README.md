# event
event-driven design in go 


```go
func exampleA() {
	e := event.New("mail.push")
	e.AddListener(func(mailto string, title string, content string) {
		fmt.Printf("mail to :%s,title:%s,content:%s\n", mailto, title, content)
	})
	e.Fire("zhengchao.deng@meican.com", "test email", "don't reply")
}
```
Things can be done easily with listener-way. 
