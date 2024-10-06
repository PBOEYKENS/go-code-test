package usecase

import (
	"context"
	"time"

	"github.com/soltix-dev/go-code-test/feature/collection/domain"
	collectionDomain "github.com/soltix-dev/go-code-test/feature/collection/domain"
)

type getCollectionHeavyUsecase struct {
	eventsRepository collectionDomain.EventsRepository
	contextTimeout   time.Duration
}

// Calls all the repository functions
// Builds the use case: necessary as functions because we are passing it to another func
func NewGetCollectionHeavyUsecase(
	eventsRepository collectionDomain.EventsRepository,
	timeout time.Duration,
) domain.GetCollectionHeavyUsecase {
	return &getCollectionHeavyUsecase{
		eventsRepository: eventsRepository,
		contextTimeout:   timeout,
	}
}

func (getCollectionHeavyUsecase *getCollectionHeavyUsecase) GetCollectionHeavyData(
	c context.Context,
	collectionId string,
	domainAddress string,
) ([]domain.EventDataModel, []domain.ReviewDataModel, error) {
	ctx, cancel := context.WithTimeout(c, getCollectionHeavyUsecase.contextTimeout)
	defer cancel()

	return getCollectionHeavyUsecase.eventsRepository.GetCollectionHeavyData(ctx, collectionId, domainAddress)
}
