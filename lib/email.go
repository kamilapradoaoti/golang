package lib

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type TipoEmail int

const (
	TE_HTML TipoEmail = iota
	TE_TEXT
)

type Email struct {
	Host        string    `json:"host"`
	Port        int       `json:"port"`
	User        string    `json:"user"`
	Pass        string    `json:"pass"`
	UseTls      bool      `json:"use_tls"`
	Subject     string    `json:"subject"`
	Attachments []string  `json:"attachments"`
	Recipients  []string  `json:"recipients"`
	Body        string    `json:"body"`
	Tipo        TipoEmail `json:"tipo"`
}

/*
GetTipo

Indica o tipo do corpo colocado no e-mail se é um texto normal ou se é HTML
*/
func (e *Email) GetTipo() string {
	switch e.Tipo {
	case TE_HTML:
		return "text/html"
	case TE_TEXT:
		return "text/plain"
	}
	return "text/plain"
}

/*
NewEmail

passa valor direto para struct
*/
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

/*
SendEmail
envia todos e-mails juntos
*/
func (e *Email) SendEmail() error {
	// Criando uma nova mensagem
	m := gomail.NewMessage()
	//quem envia
	m.SetHeader("From", e.User)
	//quem recebe
	m.SetHeader("To", e.Recipients...)
	//assunto
	m.SetHeader("Subject", e.Subject)
	//corpo do e-mail
	m.SetBody(e.GetTipo(), e.Body)

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

/*
SendEmailSingle
envia os e-mails para cada remetente separadamente
*/
func (e *Email) SendEmailSingle() error {
	//Criando uma nova mensagem
	m := gomail.NewMessage()
	//quem envia
	m.SetHeader("From", e.User)
	//assunto
	m.SetHeader("Subject", e.Subject)
	//corpo do e-mail
	m.SetBody(e.GetTipo(), e.Body)

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
	//são os que vão receber os e-mails
	m.SetHeader("To", e.Recipients...)

	return nil
}
