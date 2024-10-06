package collection

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/soltix-dev/go-code-test/core/bootstrap"
	"github.com/soltix-dev/go-code-test/feature/collection/controller"
	"github.com/soltix-dev/go-code-test/feature/collection/repository"
	"github.com/soltix-dev/go-code-test/feature/collection/usecase"
	"gorm.io/gorm"
)

func NewGetExplorePageDataRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	exploreRepo := repository.NewExploreRepository(
		db,
	)

	getExplorePageDataController := controller.GetExplorePageController{
		GetExplorePageUsecase: usecase.NewGetExplorePageDataUsecase(
			exploreRepo,
			timeout,
		),
		Env: env,
	}
	group.POST("/getExplorePageData", getExplorePageDataController.GetExplorePageData)
}
