package teste

import (
	"testing"

	"github.com/kamilapradoaoti/golang/lib"
	"github.com/stretchr/testify/assert"
)

/*
TestSendEmailConjunto
testa o envio do e-mail em conjunto
*/
func TestSendEmailConjunto(t *testing.T) {
	// caso der erro no envio, devolve uma mensagem de erro/panico
	err := lib.LoadEnv()
	if err != nil {
		panic(err.Error())
	}
	//passa para a variavel os campos do .env
	var user = lib.GetEnvString("EMAIL_USER")
	var pass = lib.GetEnvString("EMAIL_PASS")
	var host = lib.GetEnvString("EMAIL_HOST")
	var port = lib.GetEnvInt("EMAIL_PORT")
	var useTls = lib.GetEnvBool("EMAIL_USE_TLS")
	var email1 = lib.GetEnvString("EMAIL1")

	//cria um novo e-mail do zero
	email := lib.NewEmail()
	//e-mail de quem está enviando
	email.SetUser(user)
	//senha de quem envia
	email.SetPass(pass)
	//servidor do email pelo qual irá enviar o e-mail
	email.SetHost(host)
	//porta do servidor
	email.SetPort(port)
	//sistema de segurança de criptografia(se usa)
	email.UseTls = useTls
	//assunto
	email.Subject = "Teste Email"
	//corpo, podendo ser texto normal ou HTML. Há uma tratativa em email.go para reconhecer automatico
	email.Body = `
<!DOCTYPE html>
<html lang="pt-br">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Página com Fundo Preto e Imagem</title>
    <style>
        body {
            background-color: black;
            color: white;
            font-family: Arial, sans-serif;
            text-align: center;
            padding: 20px;
        }
        img {
            max-width: 100%;
            height: auto;
            margin: 20px 0;
        }
        .button {
            background-color: green;
            color: white;
            border: none;
            padding: 15px 32px;
            text-align: center;
            text-decoration: none;
            display: inline-block;
            font-size: 16px;
            margin: 20px 0;
            cursor: pointer;
            border-radius: 4px;
        }
        .button:hover {
            background-color: darkgreen;
        }
    </style>
</head>
<body>
    <h1>Bem-vindo a AOTI</h1>
    <p>Veja nossa incrível oferta abaixo e clique no botão para saber mais!</p>
    <img src="https://aoti.com.br/img/01-W1920-H1080VGradienteBlue_Black.png" alt="Imagem Descritiva">
    <br>
    <button class="button" onclick="location.href='https://aoti.com.br/#/home'">Saiba mais</button>
</body>
</html>
	`
	//aqui envia anexo no e-mail
	//email.Attachments = []string{"file1.txt", "file2.txt"}
	//indica que vai receber o e-mail, separando-os apenas por virgula
	email.Recipients = []string{email1}
	//se der erro na tentativa de enviar o e-mail, devolve um erro
	err = email.SendEmail()
	if err != nil {
		panic(err)
	}
	//afirma que não retornou nenhum erro
	assert.NoError(t, err)
}

/*
TestSendEmailSingle

testa o envio do e-mail separadamente para cada destinatario
*/
func TestSendEmailSingle(t *testing.T) {
	//da erro caso falhe a tentativa de envio
	err := lib.LoadEnv()
	if err != nil {
		panic(err.Error())
	}
	//passa os campos do arquivo .env para uma variavel
	var user = lib.GetEnvString("EMAIL_USER")
	var pass = lib.GetEnvString("EMAIL_PASS")
	var host = lib.GetEnvString("EMAIL_HOST")
	var port = lib.GetEnvInt("EMAIL_PORT")
	var useTls = lib.GetEnvBool("EMAIL_USE_TLS")
	var email1 = lib.GetEnvString("EMAIL1")
	var email2 = lib.GetEnvString("EMAIL2")

	//inicia um e-mail inteiro do zero
	email := lib.NewEmail()
	//e-mail de quem esta enviando
	email.SetUser(user)
	//senha de quem esta enviando
	email.SetPass(pass)
	//servidor do e-mail pelo qual esta enviando o e-mail
	email.SetHost(host)
	//porta do servidor do e-mail
	email.SetPort(port)
	//sistema de segurança de criptografia(se usa)
	email.UseTls = useTls
	//assunto
	email.Subject = "Teste Email"
	//corpo, pode ser feito com texto normal ou HTML, há uma tratativa em email.go
	email.Body = `
<!DOCTYPE html>
<html lang="pt-br">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Página com Fundo Preto e Imagem</title>
    <style>
        body {
            background-color: black;
            color: white;
            font-family: Arial, sans-serif;
            text-align: center;
            padding: 20px;
        }
        img {
            max-width: 100%;
            height: auto;
            margin: 20px 0;
        }
        .button {
            background-color: green;
            color: white;
            border: none;
            padding: 15px 32px;
            text-align: center;
            text-decoration: none;
            display: inline-block;
            font-size: 16px;
            margin: 20px 0;
            cursor: pointer;
            border-radius: 4px;
        }
        .button:hover {
            background-color: darkgreen;
        }
    </style>
</head>
<body>
    <h1>Bem-vindo a AOTI</h1>
    <p>Veja nossa incrível Oferta abaixo e clique no botão para saber mais!</p>
    <img src="https://aoti.com.br/img/01-W1920-H1080VGradienteBlue_Black.png" alt="Imagem Descritiva">
    <br>
    <button class="button" onclick="location.href='https://aoti.com.br/#/home'">Saiba mais</button>
</body>
</html>
	`
	//envia anexo, caso tiver
	//email.Attachments = []string{"file1.txt", "file2.txt"}
	//e-mail de quem irá receber a mensagem
	email.Recipients = []string{email1, email2}

	//caso der erro na tentativa de envio de e-mail individual, devolve um erro/panico
	err = email.SendEmailSingle()
	if err != nil {
		panic(err)
	}
	//afirma que não retornou nenhum erro
	assert.NoError(t, err)
}
