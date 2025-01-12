package main

import (
	"coin-App/route"
	"coin-App/src/controllers"
	"coin-App/src/repository"
	utils "coin-App/utils/db"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func main() {
	db, err := utils.ConnectDB()
	if err != nil {
		log.Fatalf("Error in database: %v", err.Error())
		return
	}
	repo := repository.NewRepository(db)
	log.Println("Database connected succesfully")

	router := gin.Default()
	ctrl := controllers.NewController(repo)

	route.SetupRoutes(router, ctrl)

	port := ":8000"
	log.Println("App running on port:", port)

	go func() {
		if err := router.Run(port); err != nil {
			log.Fatalf("failed to start server %v", err.Error())
		}
	}()

	//******Cron Job For Deleting Expired Coins Every 5 mins*******//
	c := cron.New()
	err = c.AddFunc("*/1 * * * *", func() {
		log.Println("Running cron job: DeleteExpiredCoins")

		ctrl.DeleteExpiredCoins(nil)

	})
	if err != nil {
		log.Fatalf("Error scheduling cron job: %v", err)
	}
	c.Start()
	select {}
}
