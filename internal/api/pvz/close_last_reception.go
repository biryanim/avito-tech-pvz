package pvz

import (
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/biryanim/avito-tech-pvz/internal/converter"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"net/http"
)

func (i *Implementation) CloseLastReception(c *gin.Context) {
	pvzIdStr := c.Param("pvzId")

	pvzId, err := uuid.Parse(pvzIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid pvzId",
		})
		return
	}

	resp, err := i.pvzService.CloseReception(c.Request.Context(), pvzId)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, model.ErrNoOpenReceptions) {
			status = http.StatusBadRequest
		}
		c.JSON(status, dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, converter.ToReceptionResponse(resp))
}
