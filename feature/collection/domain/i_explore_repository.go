package domain

import (
	"context"
)

type ItemDataModel struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`     // Decimal(10,2)
	Quantity  int     `json:"quantity"`  // INT
	Category  string  `json:"category"`  // VARCHAR(50)
	CreatedAt string  `json:"createdAt"` // DATE (ISO 8601 format)
}

type ExplorerRepository interface {
	GetExplorePageData(
		c context.Context,
		domainAddress string,
	) ([]ItemDataModel, error)

	GetSingleItemData(
		c context.Context,
		domainAddress string,
		itemName string,
	) ([]ItemDataModel, error)
}
