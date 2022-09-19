package order

import (
	"github.com/aliaydins/microservice/service.order/src/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "order service is up"})
}

func (h *Handler) getOrders(c *gin.Context) {
	buyOrders, sellOrders, err := h.service.GetOrders()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"buyOrders": buyOrders, "sellOrders": sellOrders})
}

func (h *Handler) createOrder(c *gin.Context) {

	req := new(OrderDTO)
	c.BindJSON(&req)

	reqBody := &entity.Order{
		ID:       req.ID,
		UserId:   req.UserId,
		Type:     req.Type,
		Price:    req.Price,
		Quantity: req.Quantity,
	}

	token := c.GetHeader("access_token")

	err := h.service.CreateOrder(reqBody, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order received successfully"})
}
