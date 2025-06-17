package factories

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
