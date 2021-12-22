package controllers

import (
	"fmt"
	"os"
)

var (
	// LoginProviderPort this application port listener
	LoginProviderPort string

	// LoginProviderHost is this application host (only host <host>)
	LoginProviderHost string

	// LoginProviderHostEndpoint is http complete host http://<host>
	LoginProviderHostEndpoint string

	// LoginProviderSelfPublicAPIPath is route endpoint that used for proxying request to Kratos services port
	LoginProviderSelfPublicAPIPath string

	// LoginProviderCompleteHost is complete host like http://<host>:<port>
	LoginProviderCompleteHost string

	// LoginProviderSessionsSecret is secret key for session
	LoginProviderSessionsSecret string
)

// LoadLoginProviderConfiguration for load all login provider configuration variables from .env files and define several values.
// Run this function in built-in `init` function
func LoadLoginProviderConfiguration() {
	LoginProviderPort = os.Getenv("LOGIN_PROVIDER_PORT")
	LoginProviderHost = os.Getenv("LOGIN_PROVIDER_HOST")
	LoginProviderHostEndpoint = os.Getenv("LOGIN_PROVIDER_HOST_ENDPOINT")
	LoginProviderSelfPublicAPIPath = os.Getenv("LOGIN_PROVIDER_SELF_PUBLIC_API_PATH")
	LoginProviderCompleteHost = fmt.Sprintf("%s%s", LoginProviderHostEndpoint, LoginProviderPort)
	LoginProviderSessionsSecret = os.Getenv("LOGIN_PROVIDER_SESSIONS_SECRET")
}
