package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/muratovdias/my-project/internal/models"
	"net/http"
)

func (h *Handler) createAccount(ctx *gin.Context) {
	var req models.Account

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := h.service.CreateAccount(ctx, req.ToCreateAccount())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
