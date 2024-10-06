package domain

import "context"

type GetExplorePageDataResponse struct {
	Categories []CategoryDataModel `json:"categories"`
	Topics     []TopicDataModel    `json:"topics"`
}

type GetExplorePageDataUsecase interface {
	GetExplorePageData(c context.Context, domainAddress string) ([]CategoryDataModel, []TopicDataModel, error)
}
