package auth

import (
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/biryanim/avito-tech-pvz/internal/converter"
	"github.com/gin-gonic/gin"
	"net/http"
)

//TODO: добавить middleware для правильного возвращения ошибок с кодом и сообщением + возвращать модельки

func (i *Implementation) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid request"})
		return
	}

	user, err := i.authService.Register(c.Request.Context(), converter.ToUserRegistrationModelFromRegistrationDTO(&req))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, converter.ToRegistrationRespFromUserModel(user))
}
