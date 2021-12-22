package controllers

import (
	"fmt"
	"os"

	"golang.org/x/oauth2"
)

// ORYHydraResponse struct for parse json (data response) after make request to ORY Hydra API Calls
type ORYHydraResponse struct {
	RedirectTo string `json:"redirect_to"`
}

var (
	// HydraOAuth2State is OAuth 2.0 state
	HydraOAuth2State string

	// HydraOAuth2Enpoint for OAuth 2.0 endpoint configuration
	HydraOAuth2Enpoint oauth2.Endpoint

	// HydraOAuth2AdminEndpoint is ORY Hydra administrative endpoint
	HydraOAuth2AdminEndpoint string

	// HydraOAuth2PublicEndpoint is ORY Hydra public endpoint
	HydraOAuth2PublicEndpoint string

	// HydraOAuth2RedirectPath string
	HydraOAuth2RedirectPath string

	// HydraOAuth2RedirectURL is callback url after OAuth 2.0 authentication flow finished
	HydraOAuth2RedirectURL string

	// HydraOAuth2AuthURL is ORY Hydra auth url
	HydraOAuth2AuthURL string

	// HydraOAuth2TokenURL is ORY Hydra token url
	HydraOAuth2TokenURL string
)

// LoadHydraConfiguration for load all hydra configuration variables from .env files and define several values.
// Run this function in built-in `init` function
func LoadHydraConfiguration() {
	HydraOAuth2AdminEndpoint = os.Getenv("HYDRA_OAUTH2_ADMIN_ENDPOINT")
	HydraOAuth2PublicEndpoint = os.Getenv("HYDRA_OAUTH2_PUBLIC_ENDPOINT")
	HydraOAuth2RedirectPath = os.Getenv("HYDRA_OAUTH2_REDIRECT_PATH")
	HydraOAuth2RedirectURL = fmt.Sprintf("%s%s", LoginProviderCompleteHost, HydraOAuth2RedirectPath)
	HydraOAuth2AuthURL = fmt.Sprintf("%s/oauth2/auth", HydraOAuth2AdminEndpoint)
	HydraOAuth2TokenURL = fmt.Sprintf("%s/oauth2/token", HydraOAuth2AdminEndpoint)

	HydraOAuth2Enpoint.AuthURL = HydraOAuth2AuthURL
	HydraOAuth2Enpoint.TokenURL = HydraOAuth2TokenURL

	HydraOAuth2State = generateOAuth2State()
}
