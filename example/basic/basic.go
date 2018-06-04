package main

import (
  _"fmt"
  "flag"
  _"log"
  _"github.com/BurntSushi/toml"
  _"github.com/migh/tenmeaca/mailer"
  "github.com/migh/tenmeaca/creator"
)

func main() {
  credentialsPtr := flag.String("credentials", "credentials.toml", "Credentials TOML file.")
  messagePtr := flag.String("message", "message.toml", "Message TOML file.")

  _ = credentialsPtr
  _ = messagePtr
  flag.Parse()

  templateFamily := "Not implemented"
  // // var msgCtx creator.Message
  // msgCtx := ""
  // templateName := "Not implemented"

  creator := creator.New(templateFamily)
  msg := creator.CreateMessage(/*msgCtx, templateName*/)
  _ = msg
  // if _, err := toml.DecodeFile(*messagePtr, &msgData); err != nil {
  //   log.Fatal(err)
  // }

  // mailer := mailer.New(credentials, service)
  // contact := mailer.Contact
  // mailer.Send(msg, contact)
}
