package order

import (
	"github.com/aliaydins/microservice/shared/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service   *Service
	secretKey string
}

func NewHandler(service *Service, secretKey string) *Handler {
	return &Handler{service: service, secretKey: secretKey}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()
	h.iRoutes(router)
	return router
}

func (h *Handler) iRoutes(router *gin.Engine) {
	router.Use(middleware.CORS())
	routerGroup := router.Group("/")
	routerGroup.GET("/health", h.health)
	routerGroup.GET("/", middleware.AuthMiddleware(h.secretKey), h.getOrders)
	routerGroup.POST("/", middleware.AuthMiddleware(h.secretKey), h.createOrder)

}
