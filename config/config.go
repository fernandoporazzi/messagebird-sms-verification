package config

import messagebird "github.com/messagebird/go-rest-api"

// Config is a global object that holds application level variables
var Config appConfig

type appConfig struct {
	// Messagebird client
	*messagebird.Client

	// ServerPort
	ServerPort int
}
