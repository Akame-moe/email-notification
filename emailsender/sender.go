package emailsender

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
	"gopkg.in/yaml.v2"
)

type EmailSender struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Nickname  string `yaml:"nickname"`
	Recipient string `yaml:"recipient"`
}

func NewEmailSender(configPath string) (sender *EmailSender) {
	sender = &EmailSender{}
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("failed to read config file:", err)
	}
	yaml.Unmarshal(data, sender)
	return sender
}

func (sender *EmailSender) Send(subject, content string, attachments ...interface{}) (ok bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			ok = false
			err = r.(error)
		}
	}()
	letter := email.NewEmail()

	letter.From = fmt.Sprintf("%s <%s>", sender.Nickname, sender.Username)
	letter.To = []string{sender.Recipient}
	letter.Subject = subject
	letter.Text = []byte(content)

	for i, item := range attachments {
		switch t := item.(type) {
		case string:
			if _, err := os.Stat(t); os.IsNotExist(err) {
				log.Println(t, "not exists.")
			} else {
				letter.AttachFile(t)
			}
		case []byte:
			letter.Attach(bytes.NewReader(t), fmt.Sprintf("Attachment_%d", i), "application/octet-stream")
		}

	}
	letter.SendWithTLS(fmt.Sprintf("%s:%d", sender.Host, sender.Port),
		smtp.PlainAuth("", sender.Username, sender.Password, sender.Host),
		&tls.Config{ServerName: sender.Host})
	return true, nil

}
