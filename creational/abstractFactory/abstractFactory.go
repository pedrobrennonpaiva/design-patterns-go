package abstract_factory

import (
	"design-patterns-go/creational/abstractFactory/factories"
	"fmt"
)

type IChannelFactory interface {
	MakeMessage() factories.IMessage
	SendMessage() error
}

func Run() {
	factory, err := getChannelFactory("email")
	if err != nil {
		fmt.Println(err)
	}

	message := factory.MakeMessage()
	fmt.Println(message.GetTitle())
	fmt.Println(message.GetBody())
	fmt.Println(message.GetRecipient())
	fmt.Println(message.GetSender())

	err = factory.SendMessage()
	if err != nil {
		fmt.Println(err)
	}

	factory, err = getChannelFactory("sms")
	if err != nil {
		fmt.Println(err)
	}

	message = factory.MakeMessage()
	fmt.Println(message.GetTitle())
	fmt.Println(message.GetBody())
	fmt.Println(message.GetRecipient())
	fmt.Println(message.GetSender())

	err = factory.SendMessage()
	if err != nil {
		fmt.Println(err)
	}
}

func getChannelFactory(channelType string) (IChannelFactory, error) {
	if channelType == "email" {
		return &factories.Email{}, nil
	}

	if channelType == "sms" {
		return &factories.Sms{}, nil
	}

	if channelType == "push" {
		return &factories.Push{}, nil
	}

	return nil, fmt.Errorf("invalid channel type: %s", channelType)
}
