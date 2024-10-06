package repository

import (
	"context"

	"github.com/soltix-dev/go-code-test/feature/collection/domain"
	"gorm.io/gorm"
)

type ExploreRepository struct {
	Database               *gorm.DB
	CategoriesTable        string
	TopicsTable            string
	CollectionsTable       string
	TopicsCollectionsTable string
}

func NewExploreRepository(
	db *gorm.DB,
) domain.ExplorerRepository {
	return &ExploreRepository{
		Database: db,
	}
}

func (exploreRepository *ExploreRepository) GetExplorePageData(
	c context.Context,
	domainAddress string,
) ([]domain.ItemDataModel, error) {
	db, err := exploreRepository.Database.DB()
	if err != nil {
		return nil, err
	}

	// Prepare the SQL query to select all items
	query := `SELECT id, name, price, quantity, category, created_at FROM items`

	// Create a slice to hold the item data
	var items []domain.ItemDataModel

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows and scan the data into the ItemDataModel struct
	for rows.Next() {
		var item domain.ItemDataModel
		if err := rows.Scan(&item.Id, &item.Name, &item.Price, &item.Quantity, &item.Category, &item.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
