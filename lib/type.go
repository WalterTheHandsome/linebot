package lib

type ServerCredential struct {
	LineChannelAccessToken string `yaml:"lineChannelAccessToken"` // line bot channel access Token
	LineChannelSecret      string `yaml:"lineChannelSecret"`      // line bot channel secret
	UserID                 string `yaml:"userID"`                 // not used yet
	AIAPIKey               string `yaml:"aiAPIKey"`               // the AI API key used
}
