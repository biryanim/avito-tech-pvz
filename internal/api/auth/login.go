package auth

import (
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/biryanim/avito-tech-pvz/internal/converter"
	"github.com/gin-gonic/gin"
	"net/http"
)

//TODO: добавить middleware для правильного возвращения ошибок с кодом и сообщением + возвращать модельки

func (i *Implementation) Login(ctx *gin.Context) {
	var req *dto.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := i.authService.Login(ctx, converter.ToLoginInfoFromDTO(req))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &dto.LoginResponse{
		Token: token,
	})
}
