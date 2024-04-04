package controller

import (
	"log"

	m "EksplorasiToolsAPI/Model"

	"github.com/go-gomail/gomail"
)

func sendEmail(reservation m.Reservation) {
	m := gomail.NewMessage()
	m.SetHeader("From", "jonathansena798@gmail.com")
	m.SetHeader("To", reservation.Email)
	m.SetHeader("Subject", "Reservation Confirmation")
	m.SetBody("text/html", "Thank you for your reservation, "+reservation.Name+"! Your table will be ready at "+reservation.Time+".")

	d := gomail.NewDialer("smtp.gmail.com", 587, "user", "password")

	// kirim email dalam goroutine
	go func() {
		if err := d.DialAndSend(m); err != nil {
			log.Println("Failed to send email:", err)
		} else {
			log.Println("Email sent successfully")
		}
	}()
}
