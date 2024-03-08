package controllers

import (
	"context"
	"fmt"
	"net/http"
	"server/configs"
	"server/data/responses"
	"server/models"
	"server/service"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SubscriberController struct {
	subscriberService service.SubscriberService
}

func NewSubscriberController(service service.SubscriberService) *SubscriberController {
	return &SubscriberController{
		subscriberService: service,
	}
}

var subscriberCollection *mongo.Collection = configs.GetUsersCollection(configs.DB, "subscribers")

// CreateTags		godoc
// @Summary			Subscribe
// @Description		Creating a Subscriber in MongoDB
// @Param			Subscriber body requests.GetSubscriberRequest true "subscriber"
// @Produce			application/json
// @Subscriber		subscriber
// @Success			200 {object} responses.Response{}
// @Router			/subscriber [post]
func CreateSubscriber() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var subscriber models.Subscriber
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&subscriber); err != nil {
			c.JSON(http.StatusBadRequest, responses.SubscriberResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&subscriber); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.SubscriberResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newSubscriber := models.Subscriber{
			Id:    primitive.NewObjectID(),
			Email: subscriber.Email,
		}
		result, err := subscriberCollection.InsertOne(ctx, newSubscriber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SubscriberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.SubscriberResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

// CreateTags		godoc
// @Summary			Delete Subscriber
// @Description		Unsubscribe
// @Param			Subscriber body requests.GetSubscriberRequest true "email"
// @Produce			application/json
// @Success			200 {object} responses.Response{}
// @Router			/subscriber [delete]
func DeleteSubscriber() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var subscriber models.Subscriber
		if err := c.BindJSON(&subscriber); err != nil {
			c.JSON(http.StatusBadRequest, responses.SubscriberResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if subscriber.Email == "" {
			c.JSON(http.StatusBadRequest, responses.SubscriberResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Email is required"}})
			return
		}
		filter := bson.M{"email": subscriber.Email}
		fmt.Print(filter)
		result, err := subscriberCollection.DeleteOne(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SubscriberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.SubscriberResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "No such email"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.SubscriberResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Unsub!"}},
		)
	}
}

// CreateTags		godoc
// @Summary			Get subscriber's list
// @Description		Get all current subscribers in my Database.
// @Produce			application/json
// @Success			200 {object} responses.Response{}
// @Router			/subscribers [get]
func GetAllSubscriber() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var subscribers []models.Subscriber
		defer cancel()

		results, err := subscriberCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SubscriberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleSubscriber models.Subscriber
			if err = results.Decode(&singleSubscriber); err != nil {
				c.JSON(http.StatusInternalServerError, responses.SubscriberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			subscribers = append(subscribers, singleSubscriber)
		}

		c.JSON(http.StatusOK,
			responses.SubscriberResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": subscribers}},
		)
	}
}
