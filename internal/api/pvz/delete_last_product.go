package pvz

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (i *Implementation) DeleteLastProduct(ctx *gin.Context) {
	pvzIdStr := ctx.Param("pvzId")

	pvzId, err := uuid.Parse(pvzIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = i.pvzService.DeleteLastProductInReception(ctx, pvzId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Товар удален"})
}
