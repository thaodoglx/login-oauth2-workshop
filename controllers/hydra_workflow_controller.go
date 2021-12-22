package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	hydraModels "github.com/ory/hydra-client-go/models"
	"gl.atisicloud.com/dev/sim-infinyscloud-authentication-provider/utils"
)

// getLoginRequest for fetch information on the authentication request using a REST API Call
func getLoginRequest(c *gin.Context, loginChallenge string) ([]byte, error) {
	loginRequestURL := fmt.Sprintf("%s/oauth2/auth/requests/login?login_challenge=%s", HydraOAuth2PublicEndpoint, loginChallenge)

	resp, err := utils.Fetch(c, http.MethodGet, loginRequestURL, nil, utils.ContentTypeJSON, false)
	if err != nil {
		log.Println("[!] fetch http request error ->", err.Error())
		return nil, err
	}

	return resp, nil
}

// acceptLoginRequest for accepting oauth2.0 login request
func acceptLoginRequest(c *gin.Context, loginRequest *hydraModels.LoginRequest) ([]byte, error) {
	acceptLoginRequestURL := fmt.Sprintf("%s/oauth2/auth/requests/login/accept?login_challenge=%s", HydraOAuth2PublicEndpoint, loginRequest.Challenge)

	acc := &hydraModels.AcceptLoginRequest{
		Subject:     &loginRequest.Client.ClientID,
		Remember:    true,
		RememberFor: 3600,
	}

	requestBody, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}

	resp, err := utils.Fetch(c, http.MethodPut, acceptLoginRequestURL, bytes.NewBuffer(requestBody), utils.ContentTypeJSON, false)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// getConsentRequest function for fetch information on the consent request
func getConsentRequest(c *gin.Context, consentChallenge string) ([]byte, error) {
	consentRequestURL := fmt.Sprintf("%s/oauth2/auth/requests/consent?consent_challenge=%s", HydraOAuth2PublicEndpoint, consentChallenge)

	resp, err := utils.Fetch(c, http.MethodGet, consentRequestURL, nil, utils.ContentTypeJSON, false)

	return resp, err
}

// acceptConsentRequest for accepting oauth2.0 consent request
func acceptConsentRequest(c *gin.Context, consentRequest *hydraModels.ConsentRequest) ([]byte, error) {
	acceptConsentRequestURL := fmt.Sprintf("%s/oauth2/auth/requests/consent/accept?consent_challenge=%s", HydraOAuth2PublicEndpoint, consentRequest.Challenge)

	acc := &hydraModels.AcceptConsentRequest{
		GrantScope:               consentRequest.RequestedScope,
		GrantAccessTokenAudience: consentRequest.RequestedAccessTokenAudience,
		Remember:                 true,
		RememberFor:              3600,
	}

	requestBody, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}

	resp, err := utils.Fetch(c, http.MethodPut, acceptConsentRequestURL, bytes.NewBuffer(requestBody), utils.ContentTypeJSON, false)

	return resp, err
}
