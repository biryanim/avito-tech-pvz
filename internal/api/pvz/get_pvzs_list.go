package pvz

import (
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/biryanim/avito-tech-pvz/internal/converter"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *Implementation) GetPVZs(ctx *gin.Context) {
	var pagination dto.PaginationRequest

	pagination.StartDate = ctx.Query("startDate")
	pagination.EndDate = ctx.Query("endDate")
	pagination.Page = ctx.Query("page")
	pagination.Limit = ctx.Query("limit")

	filter, err := converter.ToPaginationFilterFromPaginationRequest(&pagination)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := i.pvzService.GetListPVZs(ctx, filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, converter.ToPVZListReceptionsResponse(res))
}
