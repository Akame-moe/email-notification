package main

import (
	"github.com/akame-moe/email-notification/emailsender"
)

func main() {
	sender := emailsender.NewEmailSender("email.yml")
	sender.Send("meeting notification", "please arrive the meeting room in time.", "白鹿.jpg")
}
