package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/akame-moe/email-notification/emailsender"
)

func main() {
	var addr string
	var path string

	flag.StringVar(&addr, "addr", "0.0.0.0:4567", "the address to listen.")
	flag.StringVar(&path, "path", "/sendemail", "the route path.")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == path {
			r.ParseForm()
			subject := r.FormValue("subject")
			content := r.FormValue("content")
			if subject != "" && content != "" {
				log.Println("sending email:", subject, content)
				e := emailsender.NewEmailSender("email.yml")
				e.Send(subject, content)
			}
		} else {
			http.Redirect(w, r, "https://www.fbi.gov/tips", http.StatusFound)
		}
	})
	log.Fatal(http.ListenAndServe(addr, nil))
}
