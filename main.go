package main

import (
	"os"

	"github.com/ArturMLCKI/RestaurantMenagmentBackend/tree/main/database"
	"github.com/ArturMLCKI/RestaurantMenagmentBackend/tree/main/middleware"
	"github.com/ArturMLCKI/RestaurantMenagmentBackend/tree/main/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)

	router.Run(":" + port)
}
