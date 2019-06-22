# Executed commands

##### Create Function Code with environments variable
`aws lambda create-function --function-name DeleteMovie --zip-file fileb://./deployment.zip --runtime go1.x --handler main --role arn:aws:iam::109570655385:role/DeleteMovieRole --environment Variables={TABLE_NAME=movies} --region us-east-1`

##### Get AWS API Gateway - REST API ID

aws apigateway get-rest-apis --query "items[?name==\`MoviesAPI\`].id" --output text

##### Get AWS API Gateway - Resource ID

aws apigateway get-resources --rest-api-id API_ID --query "items[?path==\`/movies\`].id" --output text
