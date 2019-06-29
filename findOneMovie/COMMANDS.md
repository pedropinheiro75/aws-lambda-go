# Executed commands

##### Create Function Code
`aws lambda create-function --function-name FindOneMovie --zip-file fileb://./deployment.zip --runtime go1.x --handler main --role arn:aws:iam::109570655385:role/FindOneMovieRole --region us-east-1`

##### Update Function Code
`aws lambda update-function-code --function-name FindOneMovie --zip-file fileb://./deployment.zip --region us-east-1`

##### Update Function Configuration
`aws lambda update-function-configuration --function-name FindOneMovie --environment Variables={TABLE_NAME=movies} --region us-east-1`
