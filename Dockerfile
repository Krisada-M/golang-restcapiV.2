FROM golang:latest

RUN mkdir /app

ARG . /app

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV MONGOURI mongodb://Krisada3108:iWgSYobSHu9Knogx@cluster0-shard-00-00.lqb0l.mongodb.net:27017,cluster0-shard-00-01.lqb0l.mongodb.net:27017,cluster0-shard-00-02.lqb0l.mongodb.net:27017/Datainc-exp?authSource=admin&replicaSet=atlas-ffvelj-shard-0&w=majority&readPreference=primary&appname=MongoDB%20Compass&retryWrites=true&ssl=true

ENV DB_NAME db_project

ENV COLLECTION_NAME user

ENV PORT 80

RUN go build -o main

CMD [ "/app/main" ]