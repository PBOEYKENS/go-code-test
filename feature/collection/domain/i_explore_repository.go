package domain

import (
	"context"

	"github.com/gofrs/uuid"
)

const (
	CategoryTable          = "categories"
	TopicsTable            = "topics"
	TopicsCollectionsTable = "topics_collections"
	CollectionsTable       = "collections"
)

type CategoryDataModel struct {
	Id            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	ImageFilePath string    `json:"imageFilePath"`
}

type GenreDataModel struct {
	Id            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	ImageFilePath string    `json:"imageFilePath"`
}

type TopicDataModel struct {
	Id          uuid.UUID                  `json:"id"`
	Name        string                     `json:"name"`
	Collections []LightCollectionDataModel `json:"collections"`
}

type LightCollectionDataModel struct {
	Id            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	ImageFilePath string    `json:"imageFilePath"`
	IsLiked       bool      `json:"isLiked"`
	Bio           string    `json:"bio"`
}

type ExplorerRepository interface {
	GetExplorePageData(
		c context.Context,
		domainAddress string,
	) ([]CategoryDataModel, []TopicDataModel, error)

	GetCategoryPageData(
		c context.Context,
		categoryId string,
		domainAddress string,
	) ([]GenreDataModel, []TopicDataModel, error)
}
