package pvz

import (
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/biryanim/avito-tech-pvz/internal/converter"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (i *Implementation) Receptions(ctx *gin.Context) {
	var req dto.ReceptionsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
	}

	pvzId, err := uuid.Parse(req.PvzID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	resp, err := i.pvzService.CreateReception(ctx, pvzId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, converter.ToReceptionResponse(resp))
}
