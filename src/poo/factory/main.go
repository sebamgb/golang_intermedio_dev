package main

import "fmt"

// Interfaces
type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

/*end*/

// structs implement INotificationFactory
type SMSnotification struct {
}

func (SMSnotification) SendNotification() {
	fmt.Println("Sending notification via SMS")
}

func (SMSnotification) GetSender() ISender {
	return SMSnotificationSender{}
}

/*end*/

// struct implements ISender
type SMSnotificationSender struct {
}

func (SMSnotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSnotificationSender) GetSenderChannel() string {
	return "twilio"
}

/*end*/
