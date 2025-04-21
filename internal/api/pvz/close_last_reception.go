package pvz

import (
	"github.com/biryanim/avito-tech-pvz/internal/converter"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (i *Implementation) CloseLastReception(ctx *gin.Context) {
	pvzIdStr := ctx.Param("pvzId")

	pvzId, err := uuid.Parse(pvzIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := i.pvzService.CloseReception(ctx, pvzId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, converter.ToReceptionResponse(resp))
}
