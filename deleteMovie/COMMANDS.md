# Executed commands

##### Create Function Code with environments variable
`aws lambda create-function --function-name DeleteMovie --zip-file fileb://./deployment.zip --runtime go1.x --handler main --role arn:aws:iam::109570655385:role/DeleteMovieRole --environment Variables={TABLE_NAME=movies} --region us-east-1`

##### Get AWS API Gateway - REST API ID

aws apigateway get-rest-apis --query "items[?name==\`MoviesAPI\`].id" --output text

##### Get AWS API Gateway - Resource ID

aws apigateway get-resources --rest-api-id API_ID --query "items[?path==\`/movies\`].id" --output text

##### API Gateway Put Method
`aws apigateway put-method --rest-api-id 7wmejldp7k --resource-id ditz18 --http-method DELETE --authorization-type "NONE" --region us-east-1`

##### API Gateway Put Integration
`aws apigateway put-integration --rest-api-id 7wmejldp7k --resource-id ditz18 --http-method DELETE --type AWS_PROXY --integration-http-method DELETE --uri arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-1:109570655385:function:DeleteMovie/invocations --region us-east-1`

##### API Gateway - Disable response change
`aws apigateway put-method-response --rest-api-id 7wmejldp7k --resource-id ditz18 --http-method DELETE --status-code 200 --response-models '{"application/json": "Empty"}' --region us-east-1`

##### API Gateway - Redeploy API
`aws apigateway create-deployment --rest-api-id 7wmejldp7k --stage-name staging --region us-east-1`
