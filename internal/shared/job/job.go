package email

type Job struct {
	To       string
	Subject  string
	Template string
	Data     interface{}
}

var Queue = make(chan Job, 100)