package factories

import "fmt"

type Email struct {
}

type EmailMessage struct {
	Message
}

func (e *Email) MakeMessage() IMessage {
	return &EmailMessage{
		Message: Message{
			Title:     "Hello, World!",
			Body:      "This is a test message from the email",
			Recipient: "test@example.com",
			Sender:    "test@example.com",
		},
	}
}

func (e *Email) SendMessage() error {
	fmt.Println("Sending email message")
	return nil
}
