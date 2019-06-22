# Executed commands

##### Create Function Code with environments variable
`aws lambda create-function --function-name DeleteMovie --zip-file fileb://./deployment.zip --runtime go1.x --handler main --role arn:aws:iam::109570655385:role/DeleteMovieRole --environment Variables={TABLE_NAME=movies} --region us-east-1`
