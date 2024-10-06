package collection

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/soltix-dev/go-code-test/core/bootstrap"
	"github.com/soltix-dev/go-code-test/feature/collection/controller"
	collectionDomain "github.com/soltix-dev/go-code-test/feature/collection/domain"
	"github.com/soltix-dev/go-code-test/feature/collection/repository"
	"github.com/soltix-dev/go-code-test/feature/collection/usecase"
	"gorm.io/gorm"
)

func NewGetCollectionHeavyRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	eventsRepo := repository.NewEventsRepository(
		db,
		collectionDomain.EventsTable,
		collectionDomain.ReviewsTable,
		collectionDomain.SectionsTable,
	)

	getCollectionHeavyController := controller.GetCollectionHeavyController{
		GetCollectionHeavyUsecase: usecase.NewGetCollectionHeavyUsecase(
			eventsRepo,
			timeout,
		),
		Env: env,
	}
	group.POST("/getCollectionHeavy", getCollectionHeavyController.GetCollectionHeavy)
}
