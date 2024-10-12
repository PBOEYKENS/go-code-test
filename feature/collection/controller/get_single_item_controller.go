package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soltix-dev/go-code-test/core/bootstrap"
	"github.com/soltix-dev/go-code-test/core/domain"
	"github.com/soltix-dev/go-code-test/core/logger"
	collectionDomain "github.com/soltix-dev/go-code-test/feature/collection/domain"
)

type GetSingleItemController struct {
	GetExplorePageUsecase collectionDomain.GetExplorePageDataUsecase
	Env                   *bootstrap.Env
}

func (getSingleItemController *GetSingleItemController) GetSingleItem(c *gin.Context) {

	items, err := getSingleItemController.GetExplorePageUsecase.GetSingleItemData(
		c,
		getSingleItemController.Env.DomainAddress,
		c.Query("ItemName"),
	)
	if err != nil {
		logger.ErrorLog.Println(err)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	getExplorePageResponse := collectionDomain.GetExplorePageDataResponse{
		Items: items,
	}

	c.JSON(http.StatusOK, getExplorePageResponse)
}
