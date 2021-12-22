package controllers

import (
	"log"
	"text/template"

	"github.com/gin-gonic/gin"
)

// Home methods
func Home(c *gin.Context) {
	tpl := template.Must(template.ParseFiles("views/main.html"))
	err := tpl.Execute(c.Writer, nil)
	if err != nil {
		log.Println("[!] template execute error ->", err.Error())
		return
	}
}
