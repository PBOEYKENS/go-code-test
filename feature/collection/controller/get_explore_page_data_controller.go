package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soltix-dev/go-code-test/core/bootstrap"
	"github.com/soltix-dev/go-code-test/core/domain"
	"github.com/soltix-dev/go-code-test/core/logger"
	collectionDomain "github.com/soltix-dev/go-code-test/feature/collection/domain"
)

type GetExplorePageController struct {
	GetExplorePageUsecase collectionDomain.GetExplorePageDataUsecase
	Env                   *bootstrap.Env
}

func (getExplorePageController *GetExplorePageController) GetExplorePageData(c *gin.Context) {
	items, err := getExplorePageController.GetExplorePageUsecase.GetExplorePageData(
		c,
		getExplorePageController.Env.DomainAddress,
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
