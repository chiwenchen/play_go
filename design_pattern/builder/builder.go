package main

import "fmt"

type NotificationBuilder struct {
	Title      string
	Message    string
	Priority   int
	NotifyType string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) SetPriority(priority int) {
	nb.Priority = priority
}

func (nb *NotificationBuilder) SetNotifyType(notifyType string) {
	nb.NotifyType = notifyType
}

func (nb *NotificationBuilder) Build() (*Notification, error) {

	if nb.Priority > 5 {
		return nil, fmt.Errorf("Priority must below 5")
	}

	return &Notification{
		title:      nb.Title,
		message:    nb.Message,
		priority:   nb.Priority,
		notifyType: nb.NotifyType,
	}, nil
}
