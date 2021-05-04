package service

import (
	"fmt"
)

// MessageService interface for sending the messages to customers
type MessageService interface {
	SendChargeNotification(int) bool
}
// Implementation of the MessageService
type SMSService struct {}
// My Service
type MyService struct {
	MessageService MessageService
}

// implementation of sending notifications
func (sms SMSService) SendChargeNotification(value int) bool {
	fmt.Println("Sending production charge notification")
	return true
}

// ChargeCustomer - charges a customer for our service
func (ms MyService) ChargeCustomer(value int) error {
	ms.MessageService.SendChargeNotification(value) // we want to mock this SendChargeNotification
	fmt.Println("Charing the customer for value")

	return nil

}

func (ms MyService) Test(value int) error {
	ms.MessageService.SendChargeNotification(value) // we want to mock this SendChargeNotification
	fmt.Println("Charing the customer for value")
	return nil

}

