package usecase

import (
	"context"
	"time"

	"github.com/soltix-dev/go-code-test/feature/collection/domain"
	collectionDomain "github.com/soltix-dev/go-code-test/feature/collection/domain"
)

type getExplorePageDataUsecase struct {
	exploreRepository collectionDomain.ExplorerRepository
	contextTimeout    time.Duration
}

// Calls all the repository functions
// Builds the use case: necessary as functions because we are passing it to another func
func NewGetExplorePageDataUsecase(
	exploreRepository collectionDomain.ExplorerRepository,
	timeout time.Duration,
) domain.GetExplorePageDataUsecase {
	return &getExplorePageDataUsecase{
		exploreRepository: exploreRepository,
		contextTimeout:    timeout,
	}
}

func (getExplorePageData *getExplorePageDataUsecase) GetExplorePageData(
	c context.Context,
	domainAddress string,
) (
	[]collectionDomain.ItemDataModel,
	error,
) {
	ctx, cancel := context.WithTimeout(c, getExplorePageData.contextTimeout)
	defer cancel()

	return getExplorePageData.exploreRepository.GetExplorePageData(ctx, domainAddress)
}

func (getSingleItemData *getExplorePageDataUsecase) GetSingleItemData(
	c context.Context,
	domainAddress string,
	itemName string,
) (
	[]collectionDomain.ItemDataModel,
	error,
) {
	ctx, cancel := context.WithTimeout(c, getSingleItemData.contextTimeout)
	defer cancel()

	return getSingleItemData.exploreRepository.GetSingleItemData(ctx, domainAddress, itemName)
}
