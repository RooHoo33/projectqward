package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"strings"
)

type EmailResponse struct {
	Recipient string
	Message   string
	Subject   string
}

type smtpServer struct {
	host string
	port string
}

func sayHello(w http.ResponseWriter, r *http.Request) {

	var emailResponse EmailResponse

	jsonError := json.NewDecoder(r.Body).Decode(&emailResponse)
	if jsonError != nil {
		http.Error(w, jsonError.Error(), http.StatusBadRequest)
		return
	}

	from := email
	password := password

	to := []string{
		emailResponse.Recipient,
	}

	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}

	emailMessage := []byte("Subject: " + emailResponse.Subject + "\r\n" + emailResponse.Message)
	auth := smtp.PlainAuth("", from, password, smtpServer.host)

	err := smtp.SendMail(smtpServer.host+":"+smtpServer.port, auth, from, to, emailMessage)

	if err != nil {
		fmt.Println(err)
		return
	}

	message := r.URL.Path

	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message + " and name: " + emailResponse.Recipient
	//w.Write([]byte(message))
	fmt.Println(message)
}
func main() {
	http.HandleFunc("/sendemail", sayHello)
	if err := http.ListenAndServe(":4444", nil); err != nil {
		panic(err)
	}
}
