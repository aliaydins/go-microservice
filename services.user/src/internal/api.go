package user

import (
	"github.com/aliaydins/microservice/services.user/src/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user service is up"})
}

func (h *Handler) signUp(c *gin.Context) {

	req := new(entity.User)
	c.BindJSON(&req)

	user, err := h.service.SignUp(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": mapper(user)})
}
func (h *Handler) signIn(c *gin.Context) {

	req := new(entity.User)
	c.BindJSON(&req)

	user, accessToken, err := h.service.ValidateUser(req.Email, req.Password, h.secretKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrInvalidCredentials})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "user": mapper(user)})
}
