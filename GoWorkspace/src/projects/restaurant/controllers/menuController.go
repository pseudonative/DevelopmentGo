package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pseudonative/restaurant/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listening to the menu items"})
		}
		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allMenus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
