package factories

import "fmt"

type Push struct {
}

type PushMessage struct {
	Message
}

func (p *Push) MakeMessage() IMessage {
	return &PushMessage{
		Message: Message{
			Title:     "Hello, World!",
			Body:      "This is a test message from the push",
			Recipient: "device_token",
			Sender:    "app_id",
		},
	}
}

func (p *Push) SendMessage() error {
	fmt.Println("Sending push message")
	return nil
}
