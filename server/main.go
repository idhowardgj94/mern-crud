package main

import (
	"context"
	"fmt"

	"mem-crud-go/db"
	"mem-crud-go/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	// default route
	r := gin.Default()
	users := r.Group("/api/users")

	routes.UserRoute(users)

	r.GET("/test", func(c *gin.Context) {
		client, ctx, cancel := db.DbConnect()
		defer cancel()
		defer client.Disconnect(ctx)
		findUser(ctx, client)
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	r.Run()
}

var result struct {
	Name string
	Sex  string
}

func findUser(ctx context.Context, client *mongo.Client) {
	println("inside finduser")
	collection := client.Database("mern-crud").Collection("users")

	filter := bson.M{"name": "Howard Chang"}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	println("inside finduser")
	fmt.Printf("%+v\n", result)
	fmt.Printf("%s", result.Name)
}
