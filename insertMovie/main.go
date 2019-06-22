package main

import (
	"context"
	"encoding/json"
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

func insert(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var movie Movie
	if err := json.Unmarshal([]byte(request.Body), &movie); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while parsing the request body",
		},
			nil
	}

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while retrieving AWS credentials",
		},
			nil
	}

	svc := dynamodb.New(cfg)
	req := svc.PutItemRequest(&dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Item: map[string]dynamodb.AttributeValue{
			"ID": dynamodb.AttributeValue{
				S: aws.String(movie.ID),
			},
			"name": dynamodb.AttributeValue{
				S: aws.String(movie.Name),
			},
		},
	})

	if _, err := req.Send(context.Background()); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body: "Error while inserting movie to DynamoDB",
		},
		nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string("Created"),
	},
	nil
}

func main() {
	lambda.Start(insert)
}
