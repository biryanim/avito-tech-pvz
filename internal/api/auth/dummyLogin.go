package auth

import (
	"fmt"
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

	fmt.Println("1111111111111111111111")

	token, err := i.authService.DummyLogin(ctx, req.Role)
	fmt.Println("22222222222222222222")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &dto.LoginResponse{
		Token: token,
	})
}
