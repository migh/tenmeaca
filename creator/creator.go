/*
Package creator implements a library for creating HTML email.
*/

package creator

import (
  "io/ioutil"
  "github.com/matcornic/hermes"
)

type Creator struct {
  Brand Brand
}

func New(templateFamilyName string) Creator {
  return Creator{}
}

func (c *Creator) SetBrand(brand Brand) {
  c.Brand = brand
}

// Hermes Product looks more like a Brand to me
func (c *Creator) CreateMessage() string {
  h := hermes.Hermes{
    Product: c.Brand,
  }

  email := hermes.Email{
    Body: hermes.Body{
      Name: "Migsar Navarro",
      Intros: []string{
        "This is just a mail test.",
      },
      Actions: []hermes.Action{
      {
	Instructions: "Please click here:",
	Button: hermes.Button{
	  Color: "#22BC66",
	  Text: "Click me!",
	  Link: "https://www.qarmazilabs.com",
	},
      },
      },
      Outros: []string{
        "Lalala, we live like we don't care",
      },
    },
  }

  emailBody, err := h.GenerateHTML(email)
  if err != nil {
    panic(err)
  }

  err = ioutil.WriteFile("preview.html", []byte(emailBody), 0644)
  if err != nil {
    panic(err)
  }

  return emailBody
}
