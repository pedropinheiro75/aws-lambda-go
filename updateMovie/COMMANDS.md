# Executed commands

##### Create Function Code with environments variable
`aws lambda create-function --function-name UpdateMovie --zip-file fileb://./deployment.zip --runtime go1.x --handler main --role arn:aws:iam::109570655385:role/UpdateMovieTableRole --environment Variables={TABLE_NAME=movies} --region us-east-1`

##### Update function 
`aws lambda update-function-code --function-name UpdateMovie --zip-file fileb://./deployment.zip --region us-east-1 `

##### AWS API Gateway - Delete method request
`aws apigateway delete-method --rest-api-id 7wmejldp7k  --resource-id yd1262 --http-method PUT`

##### Get AWS API Gateway - REST API ID
aws apigateway get-rest-apis --query "items[?name==\`MoviesAPI\`].id" --output text

##### Get AWS API Gateway - Resource ID
aws apigateway get-resources --rest-api-id 7wmejldp7k --query "items[?path==\`/movies/{id}\`].id" --output text

##### API Gateway Put Method
`aws apigateway put-method --rest-api-id 7wmejldp7k --resource-id yd1262 --request-parameters "method.request.path.id=true" --http-method PUT --authorization-type "NONE" --region us-east-1`

##### API Gateway Put Integration
`aws apigateway put-integration --rest-api-id 7wmejldp7k --resource-id yd1262 --http-method PUT --type AWS_PROXY --integration-http-method PUT --uri arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-1:109570655385:function:UpdateMovie/invocations --region us-east-1`

##### API Gateway - Disable response change
`aws apigateway put-method-response --rest-api-id 7wmejldp7k --resource-id yd1262 --http-method PUT --status-code 200 --response-models '{"application/json": "Empty"}' --region us-east-1`

##### API Gateway - Redeploy API
`aws apigateway create-deployment --rest-api-id 7wmejldp7k --stage-name staging --region us-east-1`
