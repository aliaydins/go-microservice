package user

import (
	"github.com/aliaydins/microservice/services.user/src/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
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
	routerGroup.POST("/login", h.signIn)
	routerGroup.POST("/register", h.signUp)
}
