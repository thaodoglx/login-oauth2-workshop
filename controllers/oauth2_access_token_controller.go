package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"

	hydraModels "github.com/ory/hydra-client-go/models"
)

// GenerateAccessToken is controller for generate new access token from given client_id and client_secret
func GenerateAccessToken(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		log.Println("[!] parse form error ->", err.Error())
		return
	}

	clientID := c.PostForm("client_id")
	clientSecret := c.PostForm("client_secret")

	// store client_id and client_secret in flash session
	session, err := oauth2Store.Get(c.Request, "oauth2-flash-message")
	if err != nil {
		log.Println("[!] get session error ->", err.Error())
		return
	}

	session.AddFlash(clientID, "client_id")
	session.AddFlash(clientSecret, "client_secret")
	session.Save(c.Request, c.Writer)

	config := generateOAuth2Config(clientID, clientSecret)

	// redirect to auth code url
	authCodeURL := config.AuthCodeURL(HydraOAuth2State)

	redirectURL, err := url.Parse(authCodeURL)
	if err != nil {
		log.Println("[!] url parse error ->", err.Error())
		return
	}

	redirectURL.Host = fmt.Sprintf("%s:4444", LoginProviderHost)
	c.Redirect(http.StatusTemporaryRedirect, redirectURL.String())
}

// OAuth2Login for oauth 2.0 login controller, get login_challenge query from this controller
func OAuth2Login(c *gin.Context) {
	loginChallenge := c.Query("login_challenge")

	loginRequest, err := getLoginRequest(c, loginChallenge)
	if err != nil {
		log.Println("[!] oauth2 login request error ->", err.Error())
		return

	}
	parseLoginRequest := &hydraModels.LoginRequest{}
	err = json.Unmarshal(loginRequest, parseLoginRequest)
	if err != nil {
		log.Println("[!] json unmarshal error ->", err.Error())
		return
	}

	acc, err := acceptLoginRequest(c, parseLoginRequest)
	if err != nil {
		log.Println("[!] oauth2 accept login request error ->", err.Error())
		return
	}

	hydraResponse := &ORYHydraResponse{}
	err = json.Unmarshal(acc, hydraResponse)
	if err != nil {
		log.Println("[!] json unmarshal error ->", err.Error())
		return
	}

	redirectURL, err := url.Parse(hydraResponse.RedirectTo)
	if err != nil {
		log.Println("[!] url parse error ->", err.Error())
		return
	}

	redirectURL.Host = fmt.Sprintf("%s:4444", LoginProviderHost)
	c.Redirect(http.StatusTemporaryRedirect, redirectURL.String())
}

// OAuth2Consent for oauth 2.0 consent workflow
func OAuth2Consent(c *gin.Context) {
	consentChallenge := c.Query("consent_challenge")

	consentRequest, err := getConsentRequest(c, consentChallenge)
	if err != nil {
		log.Println("[!] oauth2 get consent request error ->", err.Error())
		return
	}

	cr := &hydraModels.ConsentRequest{}
	err = json.Unmarshal(consentRequest, cr)
	if err != nil {
		log.Println("[!] json unmarshal error ->", err.Error())
		return
	}

	acc, err := acceptConsentRequest(c, cr)
	if err != nil {
		log.Println("[!] oauth2 accept consent request ->", err.Error())
		return
	}

	hydraResponse := &ORYHydraResponse{}
	err = json.Unmarshal(acc, hydraResponse)
	if err != nil {
		log.Println("[!] json unmarshal error ->", err.Error())
		return
	}

	redirectURL, err := url.Parse(hydraResponse.RedirectTo)
	if err != nil {
		log.Println("[!] url parse error ->", err.Error())
		return
	}

	redirectURL.Host = fmt.Sprintf("%s:4444", LoginProviderHost)
	c.Redirect(http.StatusTemporaryRedirect, redirectURL.String())
}

// OAuth2Callback for oauth 2.0 callback after hydra authentication workflow finished
func OAuth2Callback(c *gin.Context) {
	session, err := oauth2Store.Get(c.Request, "oauth2-flash-message")
	if err != nil {
		log.Println("[!] get session error ->", err.Error())
		return
	}

	// get client_id from flash message
	clientID := session.Flashes("client_id")
	clientSecret := session.Flashes("client_secret")

	code := c.Query("code")
	state := c.Query("state")

	if state != HydraOAuth2State {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid oauth2 state",
		})
		return
	}

	config := generateOAuth2Config(clientID[0].(string), clientSecret[0].(string))
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Println("[!] oauth2 config exchange error ->", err.Error())
		return
	}

	fmt.Fprintf(c.Writer, "Access Token: %s\n", token.AccessToken)
	fmt.Fprintf(c.Writer, "Refresh Token: %s\n", token.RefreshToken)
	fmt.Fprintf(c.Writer, "Token Type: %s\n", token.TokenType)
	fmt.Fprintf(c.Writer, "Expiry: %s\n", token.Expiry)

	// remove sessions of flash message
	session.Options.MaxAge = -1
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		log.Println("[!] destroy flash message error ->", err.Error())
		return
	}

	return
}

func generateOAuth2Config(clientID string, clientSecret string) *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  HydraOAuth2RedirectURL,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"offline", "openid"},
		Endpoint:     HydraOAuth2Enpoint,
	}
}

func generateOAuth2State() string {
	s := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	se := base64.StdEncoding.EncodeToString(s)

	return se
}
