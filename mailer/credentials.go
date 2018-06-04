package mailer

type Credential struct {
  Name string
  Domain string
  PrivateApiKey string
  PublicValidationKey string
}

type Credentials struct {
  Credentials []Credential
}
