package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
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

func main() {
	auth = smtp.PlainAuth("", "correo", "contraseña", "smtp.gmail.com")

	templateData := struct {
		Nombre    string
		Edad      int
		Ocupacion string
	}{
		Nombre:    "Luis Arturo",
		Edad:      21,
		Ocupacion: "Estudiante",
	}

	r := NewRequest([]string{"kingarturs21@gmail.com", "lgarcia103@alumnos.uaq.mx"}, "Hola Mundo!")
	err := r.ParseTemplate(templateRoute+"index.html", templateData)

	if err != nil {
		fmt.Println("Ocurrió un error al procesar el HTML")
	}

	err = r.SendMail()
	if err != nil {
		fmt.Println("Ocurrió un error al enviar el email")
	}
}
