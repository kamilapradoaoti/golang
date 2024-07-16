package main

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func SendEmail(usuario, senha string, to, subject, body string, attachmentPath []string) error {
	from := usuario

	password := senha

	// Configurações do servidor SMTP
	smtpHost := "smtp-mail.outlook.com"
	smtpPort := 587

	// Criando uma nova mensagem
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html; charset=utf-8", body)

	// Adicionando anexo se fornecido
	if len(attachmentPath) > 0 {
		for _, nome := range attachmentPath {
			m.Attach(nome)
		}
	}

	// Configurando o dialer(cria uma discagem via e-mail) com autenticação
	d := gomail.NewDialer(smtpHost, smtpPort, from, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Enviando o e-mail
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
