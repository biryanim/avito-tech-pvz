package pvz

import (
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/biryanim/avito-tech-pvz/internal/converter"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *Implementation) CreatePvz(c *gin.Context) {
	var req dto.PVZCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request",
		})
		return
	}

	pvzInfo, err := converter.ToPVZInfoFromDTO(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	res, err := i.pvzService.CreatePVZ(c.Request.Context(), pvzInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, converter.ToPVZResponseFromPVZ(res))
}
