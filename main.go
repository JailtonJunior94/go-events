package main

import (
	"github.com/jailtonjunior94/go-events/event"
	"github.com/jailtonjunior94/go-events/user"
)

func main() {
	ed := event.NewEventDispatcher()
	sendEmailListener := user.NewSendEmailListener()
	ed.AddListener("user_created", sendEmailListener)

	user := user.NewUser(ed)
	user.Create("Jailton")
}
