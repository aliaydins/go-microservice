package history

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "history service is up"})
}

func (h *Handler) getList(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrNotValidID.Error()})
		return
	}
	history, err := h.service.GetListById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrNotFoundAnyData.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"history": history})

}
