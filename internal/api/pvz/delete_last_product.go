package pvz

import (
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"net/http"
)

func (i *Implementation) DeleteLastProduct(c *gin.Context) {
	pvzIdStr := c.Param("pvzId")

	pvzId, err := uuid.Parse(pvzIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid pvzId",
		})
		return
	}

	err = i.pvzService.DeleteLastProductInReception(c.Request.Context(), pvzId)
	if err != nil {
		if errors.Is(err, model.ErrNoSuchPvz) || errors.Is(err, model.ErrNoOpenReceptions) || errors.Is(err, model.ErrNoProducts) {
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
	c.JSON(http.StatusOK, gin.H{"message": "Товар удален"})
}
