package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// default route
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		findUser()
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	r.Run()
}

func findUser() {
	var result struct {
		Name string
		Sex  string
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("mern-crud").Collection("users")

	filter := bson.M{"name": "Howard Chang"}
	err = collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", result)
	fmt.Printf("%s", result.Name)
}
