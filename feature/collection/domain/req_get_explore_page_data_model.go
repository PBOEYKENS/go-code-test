package domain

import "context"

type GetExplorePageDataResponse struct {
	Items []ItemDataModel `json:"item"`
}

type GetExplorePageDataUsecase interface {
	GetExplorePageData(c context.Context, domainAddress string) ([]ItemDataModel, error)
	GetSingleItemData(c context.Context, domainAddress string, itemName string) ([]ItemDataModel, error)
}
