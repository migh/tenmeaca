package mailer

type Message struct {
  To []Contact
  Subject string
  Body string
  Text string
}
