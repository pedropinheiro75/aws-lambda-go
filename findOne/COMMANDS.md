# Executed commands

##### Create Function Code
`aws lambda create-function --function-name FindOneMovie --zip-file fileb://./deployment.zip --runtime go1.x --handler main --role arn:aws:iam::109570655385:role/FindOneMovieRole --region us-east-1`
