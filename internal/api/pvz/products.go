package pvz

import (
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/biryanim/avito-tech-pvz/internal/converter"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func (i *Implementation) Products(c *gin.Context) {
	var req dto.ProductCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body",
		})
		return
	}

	productPvz, err := converter.ToProductPvzFromDTO(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid product parameters",
		})
		return
	}
	resp, err := i.pvzService.AddProductToReception(c.Request.Context(), productPvz)
	if err != nil {
		if errors.Is(err, model.ErrNoOpenReceptions) {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "internal server error",
		})
		return
	}
	c.JSON(http.StatusCreated, converter.ToProductResponse(resp))
}
