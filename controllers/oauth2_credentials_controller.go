package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	hydraModels "github.com/ory/hydra-client-go/models"
	"gl.atisicloud.com/dev/sim-infinyscloud-authentication-provider/utils"
)

// Credentials controller for mana oauth2 client
func Credentials(c *gin.Context) {
	tpl := template.Must(template.ParseFiles("views/credentials.html"))
	err := tpl.Execute(c.Writer, nil)
	if err != nil {
		log.Println("[!] template execute error ->", err.Error())
		return
	}
}

// CreateOAuth2Client function for create new OAuth 2.0 client
func CreateOAuth2Client(c *gin.Context) {
	// get session:
	session, err := GetCurrentSession(c)
	if err != nil {
		log.Println("[!] get current session error ->", err.Error())
		return
	}

	// define oauth 2.0 client
	hydraOAuth2Client := hydraModels.OAuth2Client{
		RedirectUris:            []string{HydraOAuth2RedirectURL},
		GrantTypes:              []string{"authorization_code", "refresh_token"},
		ResponseTypes:           []string{"code", "id_token"},
		Scope:                   "openid offline",
		ClientName:              session.Identity.Traits.Username,
		Contacts:                []string{session.Identity.Traits.Email},
		TokenEndpointAuthMethod: "client_secret_post",
	}

	// parse struct as json for http request body
	requestBody, err := json.Marshal(hydraOAuth2Client)

	// define endpoint for hit Hydra API
	endpoint := fmt.Sprintf("%s/clients", HydraOAuth2PublicEndpoint)

	// send request to endpoint
	res, err := utils.Fetch(c, http.MethodPost, endpoint, bytes.NewBuffer(requestBody), utils.ContentTypeJSON, false)
	if err != nil {
		log.Println("[!] fetch request error ->", err.Error())
		return
	}

	// if successfully create oauth 2.0, ory hydra send response, we'll parsed the response data.
	hydraResponse := &hydraModels.OAuth2Client{}
	err = json.Unmarshal(res, hydraResponse)
	if err != nil {
		log.Println("[!] json unmarshall error ->", err.Error())
		return
	}

	fmt.Fprintf(c.Writer, "Client ID: %s\n", hydraResponse.ClientID)
	fmt.Fprintf(c.Writer, "Client Secret: %s\n", hydraResponse.ClientSecret)
}
