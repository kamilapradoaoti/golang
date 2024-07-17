package teste

import (
	"testing"

	"github.com/kamilapradoaoti/golang/lib"
	"github.com/stretchr/testify/assert"
)

func TestSendEmailConjunto(t *testing.T) {
	err := lib.LoadEnv()
	if err != nil {
		panic(err.Error())
	}
	var user = lib.GetEnvString("EMAIL_USER")
	var pass = lib.GetEnvString("EMAIL_PASS")
	var host = lib.GetEnvString("EMAIL_HOST")
	var port = lib.GetEnvInt("EMAIL_PORT")
	var useTls = lib.GetEnvBool("EMAIL_USE_TLS")

	email := lib.NewEmail()
	email.SetUser(user)
	email.SetPass(pass)
	email.SetHost(host)
	email.SetPort(port)
	email.UseTls = useTls
	email.Subject = "Test Email"
	//email.Attachments = []string{"file1.txt", "file2.txt"}
	email.Recipients = []string{"krobrelus@gmail.com", "kamilasprado@hotmail.com"}

	err = email.SendEmail()
	if err != nil {
		panic(err)
	}

	assert.NoError(t, err)
}
func TestSendEmailSingle(t *testing.T) {
	err := lib.LoadEnv()
	if err != nil {
		panic(err.Error())
	}
	var user = lib.GetEnvString("EMAIL_USER")
	var pass = lib.GetEnvString("EMAIL_PASS")
	var host = lib.GetEnvString("EMAIL_HOST")
	var port = lib.GetEnvInt("EMAIL_PORT")
	var useTls = lib.GetEnvBool("EMAIL_USE_TLS")

	email := lib.NewEmail()
	email.SetUser(user)
	email.SetPass(pass)
	email.SetHost(host)
	email.SetPort(port)
	email.UseTls = useTls
	email.Subject = "Test Email"
	//email.Attachments = []string{"file1.txt", "file2.txt"}
	email.Recipients = []string{"krobrelus@gmail.com", "kamilasprado@hotmail.com"}

	err = email.SendEmailSingle()
	if err != nil {
		panic(err)
	}

	assert.NoError(t, err)
}
