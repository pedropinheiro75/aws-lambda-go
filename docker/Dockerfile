FROM golang:latest

COPY start.sh /start.sh

RUN apt-get update && \
    apt-get install -y zip && \
    chmod +x /start.sh

EXPOSE 8080

WORKDIR /usr/src/aws-lambda-go

ENTRYPOINT [ "/start.sh" ]
