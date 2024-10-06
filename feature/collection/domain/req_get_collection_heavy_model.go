package domain

import "context"

type ReqCollectionHeavyModel struct {
	CollectionId string `form:"collectionId" json:"collectionId" binding:"required"`
}

type GetCollectionHeavyResponse struct {
	Events  []EventDataModel  `json:"events"`
	Reviews []ReviewDataModel `json:"reviews"`
}

type GetCollectionHeavyUsecase interface {
	GetCollectionHeavyData(
		c context.Context,
		collectionId string,
		domainAddress string,
	) ([]EventDataModel, []ReviewDataModel, error)
}
