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

/* ==== SMS ==== */

// structs implements INotificationFactory for SMS
type SMSnotification struct {
}

func (SMSnotification) SendNotification() {
	fmt.Println("Sending notification via SMS")
}

func (SMSnotification) GetSender() ISender {
	return SMSnotificationSender{}
}

/*end*/

// struct implements ISender for SMS
type SMSnotificationSender struct {
}

func (SMSnotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSnotificationSender) GetSenderChannel() string {
	return "twilio"
}

/*end*/
/* endSMS */

/* ==== Email ==== */

// struct implements INotification for email
type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sendng notificatio via Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

/* end */

// struct implements ISender for Email
type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

/* end */
/* endEmail */

// return structs funtion
func GetNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSnotification{}, nil
	}
	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("no notification type")
}

/* end */
// interface functions
func sendNotification(i INotificationFactory) {
	i.SendNotification()
}

func getMethod(i INotificationFactory) {
	fmt.Println(i.GetSender().GetSenderMethod())
}
func getChannel(i INotificationFactory) {
	fmt.Println(i.GetSender().GetSenderChannel())
}

/* end */
func main() {
	smsFactory, _ := GetNotificationFactory("SMS")
	emailFactory, _ := GetNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)

	getChannel(smsFactory)
	getChannel(emailFactory)
}
