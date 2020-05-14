package eye_data

import (
	"log"
	"net/mail"
	"net/smtp"

	"github.com/scorredoira/email"
)

func SendPDFtoEmail(datatime, Email, pdf_path, userID string) {
	// compose the message
	m := email.NewMessage("Результаты периметрии " + datatime, "Результаты компьютерной периметрии на " + datatime + " для UserID:" + userID)
	m.From = mail.Address{Name: "AI-Tonometry Service", Address: ""}
	m.To = []string{Email}

	// add attachments
	if err := m.Attach(pdf_path); err != nil {
		log.Panic(err)
	}

	auth := smtp.PlainAuth("", "", "", "")
	if err := email.Send("", auth, m); err != nil {
		log.Panic(err)
	}
}
