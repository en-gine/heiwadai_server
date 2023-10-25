package action

import (
	"log"
	"net/smtp"

	"server/core/infra/action"
	"server/infrastructure/env"

	"github.com/streadway/amqp"
)

var _ action.ISendMailAction = &SendMail{}

type SendMail struct{}

func (s *SendMail) SendAll(mails *[]string) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"email_queue", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to register a consumer", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			// sendMail(string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return err
}

func (s *SendMail) Send(To string, CC string, From string, Title string, Body string) error {
	msg := "From: " + From + "\n" +
		"To: " + To + "\n" +
		"Subject: " + Title + "\n\n" +
		Body

	host := env.GetEnv(env.MailHost)
	port := env.GetEnv(env.MailPort)
	pass := env.GetEnv(env.MailPass)

	err := smtp.SendMail(host+":"+port,
		smtp.PlainAuth("", From, pass, host),
		From, []string{To}, []byte(msg))
	if err != nil {
		log.Fatalf("smtp error: %s", err)
		return nil
	}

	return nil
}
