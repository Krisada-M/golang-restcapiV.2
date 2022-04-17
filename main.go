package main

import (
	"api/config"
	"api/controllers"
	"api/services"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx            context.Context
	err            error
	app            *gin.Engine
	mongoclient    *mongo.Client
	usercollection *mongo.Collection
	userservice    services.UserService
	usercontroller controllers.UserController
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI(config.EnvMongoURI())
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connected !")

	usercollection = mongoclient.Database(config.EnvDBname()).Collection(config.EnvCollectionname())
	userservice = services.NewUserService(usercollection, ctx)
	usercontroller = controllers.New(userservice)
	app = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	baseserver := app.Group("/project1")
	usercontroller.UserRoutes(baseserver)

	log.Fatal(app.Run(config.EnvPort()))
}
