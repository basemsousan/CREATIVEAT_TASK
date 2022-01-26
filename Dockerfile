# Syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY *.go ./

#Building the API 
RUN go build -o /docker-rest-emp

EXPOSE 8081

#Execution
CMD [ "/docker-rest-emp" ]
