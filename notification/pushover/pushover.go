package pushover

import (
	"fmt"

	"github.com/gregdel/pushover"
)

func Notify(message string, token string, user string) error {
	p := pushover.New(token)
	r := pushover.NewRecipient(user)
	m := pushover.NewMessage(message)
	_, err := p.SendMessage(m, r)
	if err != nil {
		return fmt.Errorf("unable to send pushover message: %v", err)
	}

	return nil
}
