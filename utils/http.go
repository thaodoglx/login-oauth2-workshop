package utils

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	// ContentTypeJSON headers
	ContentTypeJSON = map[string]string{
		"key":   "Content-Type",
		"value": "application/json",
	}
)

// Fetch function for send request (eg. GET or POST) to request url
func Fetch(c *gin.Context, method string, reqURL string, payload io.Reader, contentType map[string]string, setCookie bool) ([]byte, error) {
	cl := &http.Client{}
	req, err := http.NewRequest(method, reqURL, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set(contentType["key"], contentType["value"])

	// if setCookie is true, copy / set all cookies in current request
	if setCookie {
		cookies := c.Request.Cookies()

		for _, cookie := range cookies {
			req.AddCookie(&http.Cookie{
				Name:    cookie.Name,
				Value:   cookie.Value,
				Path:    cookie.Path,
				Expires: cookie.Expires,
			})
		}
	}

	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
