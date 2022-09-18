package wallet

import (
	"github.com/aliaydins/microservice/service.wallet/src/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "wallet service is up"})
}

func (h *Handler) GetWallet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrNotValidWalletID.Error()})
		return
	}

	wallet, err := h.service.GetWallet(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"wallet": wallet})
}

func (h *Handler) UpdateWallet(c *gin.Context) {
	req := new(entity.Wallet)
	c.BindJSON(&req)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrNotValidUserID.Error()})
		return
	}

	wallet, err := h.service.UpdateWallet(req, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"wallet": wallet})

}
