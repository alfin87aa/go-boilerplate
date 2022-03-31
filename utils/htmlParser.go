package utils

import (
	"bytes"
	"fmt"
	"text/template"
)

type SchemaHtmlRequest struct {
	To    string
	Token string
}

func ParseHtml(fileName string, data map[string]string) string {
	html, errParse := template.ParseFiles("templates/" + fileName + ".html")

	if errParse != nil {
		defer fmt.Println("parser file html failed")
	}

	body := SchemaHtmlRequest{To: data["to"], Token: data["token"]}

	buf := new(bytes.Buffer)
	errExecute := html.Execute(buf, body)

	if errExecute != nil {
		defer fmt.Println("execute html file failed")
	}

	return buf.String()
}
