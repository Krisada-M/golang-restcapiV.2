FROM golang:latest

RUN mkdir /app

ARG . /app

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV MONGOURI LINK
ENV DB_NAME db_project

ENV COLLECTION_NAME user

ENV PORT 80

RUN go build -o main

CMD [ "/app/main" ]