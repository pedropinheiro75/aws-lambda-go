# Executed commands

##### Create Function Code
`aws lambda create-function --function-name FindAllMovies --zip-file fileb://./deployment.zip --runtime go1.x --handler main --role arn:aws:iam::109570655385:role/FindAllMoviesRole --region us-east-1`

##### Update Function Code
`aws lambda update-function-code --function-name FindAllMovies --zip-file fileb://./deployment.zip --region us-east-1`

##### Add environment variable to lambda
`aws lambda update-function-configuration --function-name FindAllMovies --environment Variables={TABLE_NAME=movies} --region us-east-1`
