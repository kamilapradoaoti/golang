package lib

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type Email struct {
	Host        string   `json:"host"`
	Port        int      `json:"port"`
	User        string   `json:"user"`
	Pass        string   `json:"pass"`
	UseTls      bool     `json:"use_tls"`
	Subject     string   `json:"subject"`
	Attachments []string `json:"attachments"`
	Recipients  []string `json:"recipients"`
}

func NewEmail() *Email {
	return &Email{}
}
func (e *Email) SetUser(user string) {
	e.User = user
}
func (e *Email) SetPass(pass string) {
	e.Pass = pass
}
func (e *Email) SetHost(host string) {
	e.Host = host
}
func (e *Email) SetPort(port int) {
	e.Port = port
}
func (e *Email) SendEmail() error {
	// Criando uma nova mensagem
	m := gomail.NewMessage()
	m.SetHeader("From", e.User)
	m.SetHeader("To", e.Recipients...)
	m.SetHeader("Subject", e.Subject)

	// Adicionando anexo se fornecido
	if len(e.Attachments) > 0 {
		for _, nome := range e.Attachments {
			m.Attach(nome)
		}
	}

	// Configurando o dialer(cria uma discagem via e-mail) com autenticação
	d := gomail.NewDialer(e.Host, e.Port, e.User, e.Pass)
	if e.UseTls {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	// Enviando o e-mail
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func (e *Email) SendEmailSingle() error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.User)

	m.SetHeader("Subject", e.Subject)

	// Adicionando anexo se fornecido
	if len(e.Attachments) > 0 {
		for _, nome := range e.Attachments {
			m.Attach(nome)
		}
	}

	// Configurando o dialer(cria uma discagem via e-mail) com autenticação
	d := gomail.NewDialer(e.Host, e.Port, e.User, e.Pass)
	if e.UseTls {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}
	for _, email := range e.Recipients {
		m.SetHeader("To", email)
		// Enviando o e-mail
		if err := d.DialAndSend(m); err != nil {
			return err
		}
	}
	m.SetHeader("To", e.Recipients...)

	return nil
}
