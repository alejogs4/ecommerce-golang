package mailservice

// Mail struct to hold the email information to be sent
type Mail struct {
	From     string
	To       []string
	Subject  string
	Template string
	Body     interface{}
}

// Service interface to send emails to users
type Service interface {
	Send(mail Mail) error
}
