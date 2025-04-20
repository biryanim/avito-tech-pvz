package auth

import (
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/biryanim/avito-tech-pvz/internal/converter"
	"github.com/gin-gonic/gin"
	"net/http"
)

//TODO: добавить middleware для правильного возвращения ошибок с кодом и сообщением + возвращать модельки

func (i *Implementation) Register(ctx *gin.Context) {
	var req dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := i.authService.Register(ctx, converter.ToUserRegistrationModelFromRegistrationDTO(&req))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, converter.ToRegistrationRespFromUserModel(user))
}
