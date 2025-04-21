package auth

import (
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/biryanim/avito-tech-pvz/internal/converter"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *Implementation) DummyLogin(c *gin.Context) {
	var req dto.LoginDummyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid request"})
		return
	}

	mod, err := converter.ToRoleModel(req.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	token, err := i.authService.DummyLogin(c.Request.Context(), mod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, &dto.LoginResponse{
		Token: token,
	})
}
