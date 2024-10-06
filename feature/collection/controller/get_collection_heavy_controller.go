package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soltix-dev/go-code-test/core/bootstrap"
	"github.com/soltix-dev/go-code-test/core/domain"
	"github.com/soltix-dev/go-code-test/core/logger"
	collectionDomain "github.com/soltix-dev/go-code-test/feature/collection/domain"
)

type GetCollectionHeavyController struct {
	GetCollectionHeavyUsecase collectionDomain.GetCollectionHeavyUsecase
	Env                       *bootstrap.Env
}

func (getCollectionHeavyController *GetCollectionHeavyController) GetCollectionHeavy(c *gin.Context) {
	var request collectionDomain.ReqCollectionHeavyModel

	err := c.ShouldBind(&request)
	if err != nil {
		logger.WarningLog.Println(err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	events, reviews, err := getCollectionHeavyController.GetCollectionHeavyUsecase.GetCollectionHeavyData(
		c,
		request.CollectionId,
		getCollectionHeavyController.Env.DomainAddress,
	)
	if err != nil {
		logger.ErrorLog.Println(err)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	getCollectionHeavyResp := collectionDomain.GetCollectionHeavyResponse{
		Events:  events,
		Reviews: reviews,
	}

	c.JSON(http.StatusOK, getCollectionHeavyResp)
}
