package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

type Movie struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var movies = []Movie{
	Movie{
		ID:   1,
		Name: "Avengers",
	},
	Movie{
		ID:   2,
		Name: "Ant-Man",
	},
	Movie{
		ID:   3,
		Name: "Thor",
	},
	Movie{
		ID:   4,
		Name: "Hulk",
	},
	Movie{
		ID:   5,
		Name: "Doctor Strange",
	},
}

func insert(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var movie Movie
	if err := json.Unmarshal([]byte(req.Body), &movie); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body: "Invalid payload",
		},
		nil
	}

	movies = append(movies, movie)

	response, err := json.Marshal(movies)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body: err.Error(),
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
	lambda.Start(insert)
}
