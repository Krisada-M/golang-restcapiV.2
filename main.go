package main

import (
	"api/config"
	"api/controllers"
	"api/services"
	"context"
	"fmt"
	"log"
	"net/http"

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
	gin.SetMode(gin.ReleaseMode)
	app = gin.Default()
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello Meen"})
}

func main() {

	defer mongoclient.Disconnect(ctx)
	app.GET("/", welcome)
	baseserver := app.Group("/api")
	usercontroller.UserRoutes(baseserver)

	log.Fatal(app.Run(":" + config.EnvPort()))
}
