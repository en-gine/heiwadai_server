package messaging

import (
	"log"
	"net/smtp"
	"server/core/entity"
	"server/core/infra/action"

	"github.com/streadway/amqp"
)

var _ action.IMailAction = &SendMail{}

type SendMail struct {
}

func (s *SendMail) SendAll([]*entity.Prefecture) error {
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
			sendMail(string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return err
}

func sendMail(email string) {
	from := "your-email@example.com"
	pass := "your-password"
	to := email

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		"Hello from our Go application."

	err := smtp.SendMail("smtp.example.com:587",
		smtp.PlainAuth("", from, pass, "smtp.example.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Fatalf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://foobar.com/\n")
}
