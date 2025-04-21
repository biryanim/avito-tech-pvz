package auth

import (
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/biryanim/avito-tech-pvz/internal/converter"
	"github.com/gin-gonic/gin"
	"net/http"
)

//TODO: добавить middleware для правильного возвращения ошибок с кодом и сообщением + возвращать модельки

func (i *Implementation) Login(c *gin.Context) {
	var req *dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid request"})
		return
	}

	token, err := i.authService.Login(c.Request.Context(), converter.ToLoginInfoFromDTO(req))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &dto.LoginResponse{
		Token: token,
	})
}
