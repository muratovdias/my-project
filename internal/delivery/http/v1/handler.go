package v1

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

type Handler struct {
	service serviceI
	logger  *slog.Logger
}

func NewHandler(service serviceI, logger *slog.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/accounts", h.createAccount)

	return router
}
