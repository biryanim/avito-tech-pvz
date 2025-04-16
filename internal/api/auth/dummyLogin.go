package auth

import (
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *Implementation) DummyLogin(ctx *gin.Context) {
	var req dto.LoginDummyRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := i.authService.DummyLogin(ctx, req.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &dto.LoginResponse{
		Token: token,
	})
}
