package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"net/http"
	"os"
)

type Movie struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func updateMovie(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	movieID := request.PathParameters["id"]
	fmt.Println("pathParameter: ", movieID)

	var movie Movie
	if err := json.Unmarshal([]byte(request.Body), &movie); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Can not parse the request body",
		},
			nil
	}

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "The aws default configurations cant be loaded",
		},
			nil
	}

	svc := dynamodb.New(cfg)

	req := svc.UpdateItemRequest(&dynamodb.UpdateItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Key: map[string]dynamodb.AttributeValue{
			"ID": dynamodb.AttributeValue{
				S: aws.String(movieID),
			},
		},
		UpdateExpression: aws.String("set #name = :n"),
		ExpressionAttributeValues: map[string]dynamodb.AttributeValue{
			":n": dynamodb.AttributeValue{
				S: aws.String(movie.Name),
			},
		},
		ExpressionAttributeNames: map[string]string{
			"#name": "name",
		},
		ReturnValues: "UPDATED_NEW",
	})

	_, err = req.Send(context.Background())
	fmt.Println("Dynamo BD response")
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		},
			nil
	}

	response, err := json.Marshal(movie)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "The response can not be parsed",
		},
			nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(response),
	},
		nil
}

func main() {
	lambda.Start(updateMovie)
}
