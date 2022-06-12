package main

import (
	"fmt"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
    RequestMethod  string `json:"RequestMethod"`
    RequestBody    string `json:"RequestBody"`
    PathParameter  string `json:"PathParameter"`
    QueryParameter string `json:"QueryParameter"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(request)
	method := request.HTTPMethod
	body := request.Body
	pathParam := request.PathParameters["pathparam"]
	queryParam := request.QueryStringParameters["queryparam"]

	res := Response{
		RequestMethod: method,
		RequestBody: body,
		PathParameter: pathParam,
		QueryParameter: queryParam,
	}
	jsonBytes, _ := json.Marshal(res)
	
	return events.APIGatewayProxyResponse{
		Body: string(jsonBytes),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
