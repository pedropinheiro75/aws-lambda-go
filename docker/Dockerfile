FROM golang:latest

COPY start.sh /start.sh

RUN chmod +x /start.sh

EXPOSE 8080

WORKDIR /usr/src/go-lambda

ENTRYPOINT [ "/start.sh" ]
