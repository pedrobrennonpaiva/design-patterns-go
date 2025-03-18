package creational

import "fmt"

type IChannelFactory interface {
	MakeMessage() IMessage
	SendMessage() error
}

func GetChannelFactory(channelType string) (IChannelFactory, error) {
	if channelType == "email" {
		return &Email{}, nil
	}

	if channelType == "sms" {
		return &Sms{}, nil
	}

	if channelType == "push" {
		return &Push{}, nil
	}

	return nil, fmt.Errorf("invalid channel type: %s", channelType)
}

// IMessage
type IMessage interface {
	SetTitle(title string)
	SetBody(body string)
	SetRecipient(recipient string)
	SetSender(sender string)

	GetTitle() string
	GetBody() string
	GetRecipient() string
	GetSender() string
}

type Message struct {
	Title     string
	Body      string
	Recipient string
	Sender    string
}

func (m *Message) SetTitle(title string) {
	m.Title = title
}

func (m *Message) SetBody(body string) {
	m.Body = body
}

func (m *Message) SetRecipient(recipient string) {
	m.Recipient = recipient
}

func (m *Message) SetSender(sender string) {
	m.Sender = sender
}

func (m *Message) GetTitle() string {
	return m.Title
}

func (m *Message) GetBody() string {
	return m.Body
}

func (m *Message) GetRecipient() string {
	return m.Recipient
}

func (m *Message) GetSender() string {
	return m.Sender
}

// Email
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

// Sms
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

// Push
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
