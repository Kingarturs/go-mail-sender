package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

var (
	auth          smtp.Auth
	templateRoute = "./templates/"
)

type Request struct {
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject string) *Request {
	return &Request{
		to:      to,
		subject: subject,
	}
}


func (r *Request) SendMail() error {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "\n"
	mail := []byte(subject + mime + r.body)
	server := "smtp.gmail.com:587"

	err := smtp.SendMail(server, auth, "Arturo", r.to, mail)
	if err != nil {
		return err
	}
	return nil
}

func (r *Request) ParseTemplate(fileName string, data interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)

	if err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

func SendSimpleEmail() gin.HandlerFunc {
	return func (c *gin.Context) {
		to := c.PostForm("to")
		subject := c.PostForm("subject")
		content := c.PostForm("content")

		err := Send(to, subject, content)
		if err != nil {
			c.JSON(500, gin.H{"status": 500, "message": "Ocurrió un error"})
			return
		}
		c.JSON(200, gin.H{"status": 200, "message": "Correo enviado!"})
	}
}

func Send(to string, subject string, content string) (err error) {
	auth = smtp.PlainAuth("", "kingdeoz21@gmail.com", "cocomoco", "smtp.gmail.com")

	templateData := struct {
		Nombre    string
		Content   string
	}{
		Nombre:    "Luis Arturo",
		Content:   content,
	}

	request := NewRequest([]string{to}, subject)
	err = request.ParseTemplate(templateRoute+"mail.html", templateData)

	if err != nil {
		fmt.Println("Ocurrió un error al procesar el HTML")
		return
	}

	err = request.SendMail()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Ocurrió un error al enviar el email")
		return
	}
	return
}
