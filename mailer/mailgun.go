package mailer

import (
  "gopkg.in/mailgun/mailgun-go.v1"
)

func NewMailgunClient(credential Credential) {
  mg := mailgun.NewMailgun(credential.Domain, credential.PrivateApiKey, credential.PublicValidationKey)
  _ = mg
}

func SomeFunc() {

  // sender := "migsar@qarmazilabs.com"
  // subject := "Just a test mail"
  // body := "Test text body"
  // recipient := "migsarnavarro@gmail.com"

  // msg := mg.NewMessage(sender, subject, body, recipient)
  // msg.SetHtml(emailBody)
  //
  // res, id, err := mg.Send(msg)
  // if err != nil {
  //   panic(err)
  // }
  //
  // fmt.Printf("ID: %s Res: %s\n", id, res)
}
