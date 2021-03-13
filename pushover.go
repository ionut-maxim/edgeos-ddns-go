package main

import (
	"log"

	"github.com/gregdel/pushover"
)

func poClient(token string, user string) (*pushover.Pushover, *pushover.Recipient) {
	return pushover.New(token), pushover.NewRecipient(user)
}

func notify(message string, recipient pushover.Recipient, app pushover.Pushover) {
	notificationMessage := pushover.NewMessage(message)

	_, err := app.SendMessage(notificationMessage, &recipient)

	if err != nil {
		log.Fatal(err)
	}

}
