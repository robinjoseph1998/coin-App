package controllers

import (
	"coin-App/src/models"
	"coin-App/src/repository"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Repo repository.Repository
}

func NewController(repo repository.Repository) *Controller {
	return &Controller{
		Repo: repo}
}

func HealthPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

//*********Add Coin***********//

func (ctrl *Controller) AddCoin(c *gin.Context) {
	var request models.Coin

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "dteails": err.Error()})
		return
	}
	request.CreatedAt = time.Now()
	if request.ExpiryDate.IsZero() {
		request.ExpiryDate = time.Now().AddDate(0, 0, 1)
		// request.ExpiryDate = time.Now().Add(48 * time.Second)
	}

	if err := ctrl.Repo.CreateCoin(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create coin", "dteails": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Coin created successfully"})
}

//********Update Coin**********//

func (ctrl *Controller) UpdateCoin(c *gin.Context) {

	var request models.Coin
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}
	if request.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "coin id required"})
		return
	}

	err := ctrl.Repo.UpdateCoin(request.ID, request)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Coin not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update coin", "details": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Coin updated successfully"})

}

// ********Get Coin by Name or ID*******//

func (ctrl *Controller) GetByNameOrId(c *gin.Context) {

	var request models.Coin
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if request.ID == 0 && request.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Either 'name' or 'id' required"})
		return
	}
	var details models.Coin
	if request.ID != 0 {
		res, err := ctrl.Repo.FindById(request.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Get details by 'id'", "details": err.Error()})
			return
		}
		details = res
	}

	if request.Name != "" {
		res, err := ctrl.Repo.FindByName(request.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Get details by 'name'", "details": err.Error()})
			return
		}
		details = res
	}
	c.JSON(http.StatusFound, gin.H{"Details": details})
}

//*****List All Coins******//

func (ctrl *Controller) ListAll(c *gin.Context) {
	res, err := ctrl.Repo.ViewAllCoins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all coins", "details": err.Error()})
		return
	}
	c.JSON(http.StatusFound, gin.H{"All Coins": res})
}

// ****** Delete Expired Coins *******//

func (ctrl *Controller) DeleteExpiredCoins(c *gin.Context) {

	currentTime := time.Now()
	coins, err := ctrl.Repo.ViewAllCoins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch coins", "details": err.Error()})
		return
	}

	for _, coin := range coins {
		if coin.ExpiryDate.Before(currentTime) {
			err := ctrl.Repo.DeleteCoin(coin.ID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete expired coin", "details": err.Error()})
				return
			}
			err = ctrl.Repo.LogExpiredCoins(coin.Name, coin.ExpiryDate)
			if err != nil {
				log.Fatalf("failed to log expired coin: %v", err)
				return
			}
		}
	}
	log.Println("expired coin deleted successfully")

}

// ******* List Expired Coins Log ********//

func (ctrl *Controller) ViewExpiredLogs(c *gin.Context) {
	res, err := ctrl.Repo.ViewExpiredCoins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch expired coin logs", "details": err.Error()})
		return
	}
	c.JSON(http.StatusFound, gin.H{"Details": res})
}
