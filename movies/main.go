package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"net/http"
	"os"
)

type Movie struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func findOne(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["id"]

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while retrieving AWS credentials",
		},
			nil
	}

	svc := dynamodb.New(cfg)
	req := svc.GetItemRequest(&dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Key: map[string]dynamodb.AttributeValue{
			"ID": dynamodb.AttributeValue{
				S: aws.String(id),
			},
		},
	})

	res, err := req.Send(context.Background())
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while fetching movie from DynamoDB",
		},
			nil
	}

	movie := Movie{
		ID:   *res.Item["ID"].S,
		Name: *res.Item["name"].S,
	}

	response, err := json.Marshal(movie)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while decoding to string value",
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

func findAll() (events.APIGatewayProxyResponse, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while retrieving AWS credentials",
		},
			nil
	}

	svc := dynamodb.New(cfg)
	req := svc.ScanRequest(&dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	})
	res, err := req.Send(context.Background())
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while scanning DynamoDB",
		},
			nil
	}

	movies := make([]Movie, 0)
	for _, item := range res.Items {
		movies = append(movies, Movie{
			ID:   *item["ID"].S,
			Name: *item["name"].S,
		})
	}

	response, err := json.Marshal(movies)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while decoding to string value",
		},
			nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-type": "application/json",
		},
		Body: string(response),
	},
		nil
}

func insertMovie(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var movie Movie

	if err := json.Unmarshal([]byte(request.Body), &movie); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid payload",
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

	if _, err = req.Send(context.Background()); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while inserting movie to DynamoDB",
		},
			nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	},
		nil
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

func deleteMovie(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var movie Movie

	if err := json.Unmarshal([]byte(request.Body), &movie); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body: "Invalid Payload",
		},
			nil
	}

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body: "Error while retrieving AWS credentials",
		},
			nil
	}

	svc := dynamodb.New(cfg)
	req := svc.DeleteItemRequest(&dynamodb.DeleteItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Key: map[string]dynamodb.AttributeValue{
			"ID": dynamodb.AttributeValue{
				S: aws.String(movie.ID),
			},
		},
	})

	if _, err := req.Send(context.Background()); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body: "Error while deleting movie from DynamoDB",
		},
			nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	},
		nil
}

func movies(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch method := request.HTTPMethod; method {
	case http.MethodGet:
		if _, ok := request.PathParameters["id"]; ok {
			return findOne(request)
		}
		return findAll()
	case http.MethodPost:
		return insertMovie(request)
	case http.MethodPut:
		return updateMovie(request)
	case http.MethodDelete:
		return deleteMovie(request)
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body: "Unsupported HTTP method",
		},
			nil
	}
}

func main() {
	lambda.Start(movies)
}
