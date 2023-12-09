package main

import (
	"bytes"
	"embed"
	"fmt"
	"net/http"
	"text/template"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//go:embed views/homepage.html
var templates embed.FS

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	data := struct {
		Title   string
		Content string
	}{
		Title:   "My Page",
		Content: "Hello, Golang!",
	}
	// fmt.Printf("Output: %+v", request)
	var tmpl *template.Template
	var err error

	if request.Path == "/" {
		tmpl, err = template.ParseFS(templates, "views/homepage.html")
	}

	var result bytes.Buffer
	if err != nil {
		fmt.Fprintf(&result, "Error parsing html files: %v , status %d", err.Error(), http.StatusInternalServerError)
		return events.APIGatewayProxyResponse{Body: result.String(), StatusCode: http.StatusInternalServerError}, err
	}

	if err := tmpl.Execute(&result, data); err != nil {
		fmt.Fprintf(&result, "Error parsing html template: %v", err)
		return events.APIGatewayProxyResponse{Body: result.String(), StatusCode: http.StatusInternalServerError}, err
	}
	return events.APIGatewayProxyResponse{
		Body:       "Hello World",
		Headers:    map[string]string{"Content-Type": "text/html"},
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
