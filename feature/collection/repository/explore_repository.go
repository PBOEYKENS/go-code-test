package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gofrs/uuid"
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
	categoryTable string,
	topicsTable string,
	collectionsTable string,
	topicsCollectionsTable string,
) domain.ExplorerRepository {
	return &ExploreRepository{
		Database:               db,
		CategoriesTable:        categoryTable,
		TopicsTable:            topicsTable,
		CollectionsTable:       collectionsTable,
		TopicsCollectionsTable: topicsCollectionsTable,
	}
}

func getTopics(
	db *sql.DB,
	query string,
	domainAddress string,
) ([]domain.TopicDataModel, error) {
	topicRows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	var topics []domain.TopicDataModel
	topicMap := make(map[uuid.UUID]*domain.TopicDataModel)

	for topicRows.Next() {
		var row struct {
			TopicId                 uuid.UUID
			TopicName               string
			CollectionId            uuid.UUID
			CollectionName          string
			CollectionImageFilePath string
			CollectionIsLiked       bool
			CollectionBio           string
		}

		if err := topicRows.Scan(
			&row.TopicId,
			&row.TopicName,
			&row.CollectionId,
			&row.CollectionName,
			&row.CollectionImageFilePath,
			&row.CollectionIsLiked,
			&row.CollectionBio,
		); err != nil {
			return nil, err
		}

		// Check if we already have this topic
		topic, exists := topicMap[row.TopicId]
		if !exists {
			// Create new topic if it doesn't exist
			topic = &domain.TopicDataModel{
				Id:          row.TopicId,
				Name:        row.TopicName,
				Collections: []domain.LightCollectionDataModel{},
			}
			topicMap[row.TopicId] = topic
		}

		// Add collection to the topic
		collection := domain.LightCollectionDataModel{
			Id:            row.CollectionId,
			Name:          row.CollectionName,
			ImageFilePath: domainAddress + row.CollectionImageFilePath,
			IsLiked:       row.CollectionIsLiked,
			Bio:           row.CollectionBio,
		}
		topic.Collections = append(topic.Collections, collection)
	}

	// Convert map to slice
	for _, topic := range topicMap {
		topics = append(topics, *topic)
	}

	if err := topicRows.Err(); err != nil {
		return nil, err
	}

	return topics, nil
}

func (exploreRepository *ExploreRepository) GetCategoryPageData(
	c context.Context,
	categoryId string,
	domainAddress string,
) ([]domain.GenreDataModel, []domain.TopicDataModel, error) {
	db, err := exploreRepository.Database.DB()
	if err != nil {
		return nil, nil, err
	}

	// Todo Get the Genre data
	genres := []domain.GenreDataModel{}

	topicsQuery := fmt.Sprintf(`
		SELECT 
			t.id AS topic_id,
			t.name AS topic_name,
			co.id AS collection_id,
			co.name AS collection_name,
			co.image_file_path AS collection_image_file_path,
			co.is_liked AS collection_is_liked,
			co.bio AS collection_bio
		FROM 
			topics t
		JOIN 
			topics_collections tc ON t.id = tc.topic_id
		JOIN 
			collections co ON tc.collection_id = co.id
		WHERE 
			t.category_id = '%v';
	`, categoryId)

	topics, err := getTopics(db, topicsQuery, domainAddress)
	if err != nil {
		return nil, nil, err
	}

	return genres, topics, nil
}

func (exploreRepository *ExploreRepository) GetExplorePageData(
	c context.Context,
	domainAddress string,
) ([]domain.CategoryDataModel, []domain.TopicDataModel, error) {
	db, err := exploreRepository.Database.DB()
	if err != nil {
		return nil, nil, err
	}

	categoryRows, err := db.Query(`
        SELECT 
            cat.id,
            cat.name,
            cat.image_file_path,
            cat.any_category
        FROM
            categories cat
    `)

	var categories []domain.CategoryDataModel
	anyCategoryId := ""
	for categoryRows.Next() {
		var row struct {
			CategoryId            uuid.UUID
			CategoryName          string
			CategoryImageFilePath string
			CategoryAnyCategory   bool
		}

		if err := categoryRows.Scan(&row.CategoryId,
			&row.CategoryName,
			&row.CategoryImageFilePath,
			&row.CategoryAnyCategory,
		); err != nil {
			return nil, nil, err
		}

		if !row.CategoryAnyCategory {
			category := domain.CategoryDataModel{
				Id:            row.CategoryId,
				Name:          row.CategoryName,
				ImageFilePath: domainAddress + row.CategoryImageFilePath,
			}
			categories = append(categories, category)
		} else {
			anyCategoryId = row.CategoryId.String()
		}
	}

	topicsQuery := fmt.Sprintf(`
		SELECT 
			t.id AS topic_id,
			t.name AS topic_name,
			co.id AS collection_id,
			co.name AS collection_name,
			co.image_file_path AS collection_image_file_path,
			co.is_liked AS collection_is_liked,
			co.bio AS collection_bio
		FROM 
			topics t
		JOIN 
			topics_collections tc ON t.id = tc.topic_id
		JOIN 
			collections co ON tc.collection_id = co.id
		WHERE 
			t.category_id = '%v';
	`, anyCategoryId)

	topics, err := getTopics(db, topicsQuery, domainAddress)
	if err != nil {
		return nil, nil, err
	}

	return categories, topics, nil
}
