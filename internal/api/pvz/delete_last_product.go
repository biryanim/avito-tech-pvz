package pvz

import (
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (i *Implementation) DeleteLastProduct(ctx *gin.Context) {
	var req dto.DeleteProductRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pvzId, err := uuid.Parse(req.PvzID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = i.pvzService.DeleteLastProductInReception(ctx, pvzId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Товар удален"})
}
