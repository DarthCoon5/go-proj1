package shop

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/internal/sell-service/application"
	"shop/internal/sell-service/application/command_workers"
	"shop/internal/sell-service/application/query_workers"
	"strconv"
)

type ProductWorkers struct {
	ShopWorker    application.ShopWorker
	QueryWorker   query_workers.ShopReceiver
	CommandWorker command_workers.ShopPusher
}

type ProductCreateRequest struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

func RegisterProductRoutes(workers *ProductWorkers, r *gin.RouterGroup) {
	r.POST("/product", workers.CreateProduct)
	r.GET("/products", workers.GetProductsList)
	r.GET("/products/count", workers.CountProducts)
}

func (worker *ProductWorkers) GetProductsList(c *gin.Context) {

	respond, err := worker.QueryWorker.ShopReceiverInterface.GetProductsList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, respond)
}

func (worker *ProductWorkers) CreateProduct(c *gin.Context) {
	var rootRequest ProductCreateRequest
	if err := c.ShouldBindJSON(&rootRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	trueRequest := command_workers.ProductRequest{
		Name:   rootRequest.Name,
		Number: rootRequest.Number,
	}

	returnList, err := strconv.ParseBool(c.Param("returnList"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if returnList {
		respond, err := worker.ShopWorker.ShopWorkerInterface.CreateProductAndList(trueRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, respond)
	} else {
		err := worker.CommandWorker.CreateProduct(trueRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, http.StatusOK)
	}
}

func (worker *ProductWorkers) CountProducts(c *gin.Context) {
	color := c.Param("color")

	if color == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no color"})
		return
	}

	respond, err := worker.QueryWorker.GetProductsCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, respond)
}
