package routes

import (
	"mem-crud-go/db"
	"mem-crud-go/models"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
)

// UserRoute user's route
func UserRoute(users *gin.RouterGroup) {
	users.GET("/:id", func(c *gin.Context) {
		client, ctx, cancel := db.Connect()
		defer cancel()
		defer client.Disconnect(ctx)

		var id, _ = primitive.ObjectIDFromHex(c.Param("id"))
		collection := client.Database("mern-crud").Collection("users")
		filter := bson.M{
			"_id": id,
		}
		sr := collection.FindOne(ctx, filter)

		if sr == nil {
			println("no value found")
			c.JSON(200, models.Response{
				Success: false,
				Msg:     `no match user.`,
			})
		} else {
			var result models.User
			sr.Decode(&result)
			c.JSON(200, &result)
		}
	})

	users.GET("/", func(c *gin.Context) {
		client, ctx, cancel := db.Connect()
		defer cancel()
		defer client.Disconnect(ctx)

		collection := client.Database("mern-crud").Collection("users")
		cur, _ := collection.Find(ctx, bson.D{})

		var results []models.User = make([]models.User, 0)
		cur.All(ctx, &results)
		c.JSON(200, &results)
	})

	users.POST("/", func(c *gin.Context) {
		client, ctx, cancel := db.Connect()
		defer cancel()
		defer client.Disconnect(ctx)

		var data = new(models.User)
		c.ShouldBindJSON(data)

		errs := data.Validate()
		if len(errs) != 0 {
			c.JSON(400, models.Response{
				Success: false,
				Msg:     &errs,
			})
			return
		}

		collection := client.Database("mern-crud").Collection("users")
		result, _ := collection.InsertOne(ctx, &data)
		sr := collection.FindOne(ctx, bson.M{
			"_id": result.InsertedID,
		})
		data = nil
		sr.Decode(&data)
		c.JSON(200, models.Response{
			Success: true,
			Msg:     "Successfully added!",
			Result:  &data,
		})
	})

	users.PUT("/:id", func(c *gin.Context) {
		var id, _ = primitive.ObjectIDFromHex(c.Param("id"))
		client, ctx, cancel := db.Connect()
		defer cancel()
		defer client.Disconnect(ctx)
		collection := client.Database("mern-crud").Collection("users")
		var input models.User
		c.ShouldBindJSON(&input)

		errs := input.Validate()
		if len(errs) != 0 {
			c.JSON(400, models.Response{
				Success: false,
				Msg:     &errs,
			})
			return
		}

		update := bson.M{
			"$set": input,
		}
		filter := bson.M{
			"_id": id,
		}
		collection.UpdateOne(ctx, filter, update)

		r := collection.FindOne(ctx, bson.M{"_id": id})
		update = nil
		r.Decode(&update)
		c.JSON(200, models.Response{
			Success: true,
			Msg:     "successfully updated!",
			Result:  &update,
		})
	})

	users.DELETE("/:id", func(c *gin.Context) {
		client, ctx, cancel := db.Connect()
		defer cancel()
		defer client.Disconnect(ctx)
		collection := client.Database("mern-crud").Collection("users")
		var id, _ = primitive.ObjectIDFromHex(c.Param("id"))

		collection.DeleteOne(ctx, bson.M{
			"_id": id,
		})

		var result struct {
			ID string `json:"_id" bson:"_id"`
		}

		result.ID = c.Param("id")
		c.JSON(200, models.Response{
			Success: true,
			Msg:     "it has been deleted",
			Result:  &result,
		})
	})
}
