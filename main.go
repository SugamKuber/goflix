package main

import (
	"context"
	"fmt"
	"log"

	"example.com/goflix/controllers"
	"example.com/goflix/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server          *gin.Engine
	movieservice    services.MovieServices
	moviecontroller controllers.MovieController
	ctx             context.Context
	moviec           *mongo.Collection
	mongoclient     *mongo.Client
	err             error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb+srv://admin:1234@cluster0.ahqouqu.mongodb.net/?retryWrites=true&w=majority")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	moviec = mongoclient.Database("moviedb").Collection("movies")
	movieservice = services.NewMovieService(moviec, ctx)
	moviecontroller = controllers.New(movieservice)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	moviecontroller.RegisterMovieRoutes(basepath)

	log.Fatal(server.Run(":8080"))

}
