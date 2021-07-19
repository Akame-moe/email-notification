package main

import (
	"flag"
	"log"
	"strings"

	"github.com/akame-moe/email-notification/emailsender"
)

type Attachments []string

func (a *Attachments) Set(val string) error {
	if val != "" {
		*a = append(*a, val)
	}
	return nil
}
func (a *Attachments) String() string {
	var r []string
	for _, t := range *a {
		r = append(r, t)
	}
	return "[" + strings.Join(r, ", ") + "]"
}

func main() {
	var subject string
	var content string
	var attachments Attachments
	flag.StringVar(&subject, "s", "", "the email subject.")
	flag.StringVar(&content, "c", "", "the email content.")
	flag.Var(&attachments, "a", "the attachments.")
	flag.Parse()
	email := emailsender.NewEmailSender("~/.email.yml")
	ok, err := email.Send(subject, content, attachments)
	if !ok {
		log.Println("error:", err)
	}
}
