# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

RUN mkdir /controllers
COPY controllers/ /controllers
RUN mkdir /models
COPY models/ models/
RUN mkdir /routes
COPY routes/ /routes

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /api-test

EXPOSE 8000

CMD [ "/api-test" ]