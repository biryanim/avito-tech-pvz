package pvz

import (
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/biryanim/avito-tech-pvz/internal/converter"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *Implementation) GetPVZs(c *gin.Context) {
	var pagination dto.PaginationRequest

	pagination.StartDate = c.Query("startDate")
	pagination.EndDate = c.Query("endDate")
	pagination.Page = c.Query("page")
	pagination.Limit = c.Query("limit")

	filter, err := converter.ToPaginationFilterFromPaginationRequest(&pagination)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := i.pvzService.GetListPVZs(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, converter.ToPVZListReceptionsResponse(res))
}
