package shop

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/internal/sell-service/application/command_workers"
	"shop/internal/sell-service/application/query_workers"
	"shop/internal/sell-service/domain"
	"shop/pkg/math"
	"strconv"
)

type ShopWorkers struct {
	QueryReceiver   query_workers.ShopReceiver
	CommandReceiver command_workers.ShopPusher
}

func RegisterShopReceiverRoutes(workers *ShopWorkers, r *gin.RouterGroup) {
	r.GET("/double_num", workers.GetDoubleNumber)
	r.GET("/time", workers.CalculateTime)
}

func (worker *ShopWorkers) GetDoubleNumber(c *gin.Context) {
	num, err := strconv.Atoi(c.Query("num"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	respond := math.GetDoubleNumber(num)
	c.JSON(http.StatusOK, respond)
}

func (worker *ShopWorkers) CalculateTime(c *gin.Context) {
	role := domain.Role(0)

	if role.IsAdmin() {
		respond := domain.CalculateAverageWaitingTime()
		c.JSON(http.StatusOK, respond)
		return
	}

	c.JSON(http.StatusForbidden, 0)
}
