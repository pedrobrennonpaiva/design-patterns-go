package factories

import "fmt"

type Sms struct {
}

type SmsMessage struct {
	Message
}

func (s *Sms) MakeMessage() IMessage {
	return &SmsMessage{
		Message: Message{
			Title:     "Hello, World!",
			Body:      "This is a test message from the sms",
			Recipient: "+1234567890",
			Sender:    "+0987654321",
		},
	}
}

func (s *Sms) SendMessage() error {
	fmt.Println("Sending sms message")
	return nil
}
