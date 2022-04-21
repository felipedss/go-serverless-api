package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func router(req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {

	if req.RequestContext.HTTP.Path == "/users" && req.RequestContext.HTTP.Method == "GET" {
		return handleGetUsers(req)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Body:       http.StatusText(http.StatusMethodNotAllowed),
	}, nil
}

func handleGetUsers(req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {

	users := []User{
		{
			ID:   1,
			Name: "João",
		},
		{
			ID:   2,
			Name: "Pedro",
		},
	}

	js, err := json.Marshal(users)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

func main() {
	lambda.Start(router)
}
