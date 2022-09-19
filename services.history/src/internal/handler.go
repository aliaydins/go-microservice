package history

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
	h.initRoutes(router)
	return router
}

func (h *Handler) initRoutes(router *gin.Engine) {
	router.Use(middleware.CORS())
	routerGroup := router.Group("/")
	routerGroup.GET("/health", h.health)
	routerGroup.GET("/:id", middleware.AuthMiddleware(h.secretKey), h.getList)
}
