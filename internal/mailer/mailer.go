package mailer

import (
	"bytes"
	"embed"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gopkg.in/gomail.v2"
)

//go:embed templates
var templateFS embed.FS

type Mailer struct {
	dialer *gomail.Dialer
	sender string // no-reply@budgetapi.com
	Logger echo.Logger
}

type EmailData struct {
	AppName  string
	Subject  string
	MetaData interface{}
}

func NewMailer(logger echo.Logger) Mailer {
	err := godotenv.Load()
	if err != nil {
		panic("error loading mail env variables")
	}

	mailHost := os.Getenv("MAIL_HOST")
	mailPort, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		logger.Fatal(err)
	}

	mailUsername := os.Getenv("MAIL_USERNAME")
	mailPassword := os.Getenv("MAIL_PASSWORD")
	mailSender := os.Getenv("MAIL_SENDER")

	d := gomail.NewDialer(mailHost, mailPort, mailUsername, mailPassword)

	return Mailer{
		sender: mailSender,
		dialer: d,
		Logger: logger,
	}
}

func (m *Mailer) Send(recipient string, templateFile string, data EmailData) error {
	absolutePath := filepath.Join("templates", templateFile)
	tmpl, err := template.ParseFS(templateFS, absolutePath)
	if err != nil {
		m.Logger.Error(err)
		return err
	}

	data.AppName = os.Getenv("APP_NAME")

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		m.Logger.Error(err)
		return err
	}

	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	if err != nil {
		m.Logger.Error(err)
		return nil
	}

	gomailMsg := gomail.NewMessage()
	gomailMsg.SetHeader("To", recipient)
	gomailMsg.SetHeader("From", m.sender)
	gomailMsg.SetHeader("Subject", subject.String())

	gomailMsg.SetBody("text/html", htmlBody.String())

	err = m.dialer.DialAndSend(gomailMsg)
	if err != nil {
		m.Logger.Error(err)
		return err
	}

	return nil
}
