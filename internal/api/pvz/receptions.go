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

func (i *Implementation) Receptions(c *gin.Context) {
	var req dto.ReceptionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid request"})
		return
	}

	pvzId, err := uuid.Parse(req.PvzID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid pvzId"})
		return
	}

	resp, err := i.pvzService.CreateReception(c.Request.Context(), pvzId)
	if err != nil {
		if errors.Is(err, model.ErrLastReceptionNotClosed) {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "internal server error"})
		return
	}
	c.JSON(http.StatusCreated, converter.ToReceptionResponse(resp))
}
