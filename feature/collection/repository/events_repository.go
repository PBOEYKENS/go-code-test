package repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"

	"github.com/gofrs/uuid"
	"github.com/lib/pq"
	"github.com/soltix-dev/go-code-test/feature/collection/domain"
)

type EventsRepository struct {
	Database      *gorm.DB
	EventsTable   string
	ReviewsTable  string
	SectionsTable string
}

func NewEventsRepository(
	db *gorm.DB,
	eventsTable string,
	reviewsTable string,
	sectionsTable string,
) domain.EventsRepository {
	return &EventsRepository{
		Database:      db,
		EventsTable:   eventsTable,
		ReviewsTable:  reviewsTable,
		SectionsTable: sectionsTable,
	}
}

func (eventsRepository *EventsRepository) GetCollectionHeavyData(
	c context.Context,
	collectionId string,
	domainAddress string,
) ([]domain.EventDataModel, []domain.ReviewDataModel, error) {
	db, err := eventsRepository.Database.DB()
	if err != nil {
		return nil, nil, err
	}

	eventsQuery := fmt.Sprintf(`
		SELECT 
			events.id, 
			events.is_single_section, 
			events.is_increment_ticket, 
			events.name, 
			events.performer, 
			EXTRACT(EPOCH FROM events.event_from) AS event_from, 
			EXTRACT(EPOCH FROM events.event_to) AS event_to, 
			events.venue, 
			events.location,
			events.image_path,
			events.event_tax,
			events.platform_tax,
			events.platform_percentage,
			events.platform_fixed,
			events.processing_percentage,
			events.processing_fixed,
			sections.id, 
			sections.seating_map, 
			sections.name, 
			sections.lowest_price, 
			sections.highest_price, 
			sections.seat_available,
			sections.seat_type_array
		FROM 
			events
		JOIN 
			collections_events ON events.id = collections_events.event_id
		LEFT JOIN 
			sections ON events.section_id = sections.id AND events.is_single_section = TRUE
		WHERE 
			collections_events.collection_id = '%v'
		ORDER BY
			events.event_from ASC;
	`, collectionId)

	eventsRows, err := db.Query(eventsQuery)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	var events []domain.EventDataModel
	for eventsRows.Next() {
		var row struct {
			EventId                *uuid.UUID
			EventIsSingleSection   *bool
			EventIsIncrementTicket *bool
			EventName              *string
			EventPerformer         *string
			EventFrom              *string
			EventTo                *string
			EventVenue             *string
			EventLocation          *string
			EventImagePath         *string
			EventTax               *float64
			PlatformTax            *float64
			PlatformPercentage     *float64
			PlatformFixed          *float64
			ProcessingPercentage   *float64
			ProcessingFixed        *float64
			SectionSeatingMap      *string
			SectionId              *uuid.UUID
			SectionName            *string
			SectionLowestPrice     *float32
			SectionHighestPrice    *float32
			SectionSeatAvailable   *bool
			SectionSeatTypes       []string
		}

		if err := eventsRows.Scan(
			&row.EventId,
			&row.EventIsSingleSection,
			&row.EventIsIncrementTicket,
			&row.EventName,
			&row.EventPerformer,
			&row.EventFrom,
			&row.EventTo,
			&row.EventVenue,
			&row.EventLocation,
			&row.EventImagePath,
			&row.EventTax,
			&row.PlatformTax,
			&row.PlatformPercentage,
			&row.PlatformFixed,
			&row.ProcessingPercentage,
			&row.ProcessingFixed,
			&row.SectionId,
			&row.SectionSeatingMap,
			&row.SectionName,
			&row.SectionLowestPrice,
			&row.SectionHighestPrice,
			&row.SectionSeatAvailable,
			pq.Array(&row.SectionSeatTypes),
		); err != nil {
			return nil, nil, err
		}

		eventFromStr := strings.Split(*row.EventFrom, ".")[0]
		eventFromInt, err := strconv.ParseUint(eventFromStr, 10, 64)
		if err != nil {
			return nil, nil, err
		}

		eventToStr := strings.Split(*row.EventTo, ".")[0]
		eventToInt, err := strconv.ParseUint(eventToStr, 10, 64)
		if err != nil {
			return nil, nil, err
		}

		event := domain.EventDataModel{
			Id:                   *row.EventId,
			IsSingleSectionEvent: *row.EventIsSingleSection,
			IsIncrementTicket:    *row.EventIsIncrementTicket,
			Name:                 *row.EventName,
			Performer:            *row.EventPerformer,
			EventFrom:            eventFromInt,
			EventTo:              eventToInt,
			Venue:                *row.EventVenue,
			Location:             *row.EventLocation,
			ImagePath:            domainAddress + *row.EventImagePath,
			EventTax:             *row.EventTax,
			PlatformTax:          *row.PlatformTax,
			PlatformPercentage:   *row.PlatformPercentage,
			PlatformFixed:        *row.PlatformFixed,
			ProcessingPercentage: *row.ProcessingPercentage,
			ProcessingFixed:      *row.ProcessingFixed,
		}

		if event.IsSingleSectionEvent {
			event.SectionModel = domain.SectionDataModel{
				Id:             *row.SectionId,
				SeatingMap:     *row.SectionSeatingMap,
				Name:           *row.SectionName,
				LowestPrice:    *row.SectionLowestPrice,
				HighestPrice:   *row.SectionHighestPrice,
				SeatTypes:      row.SectionSeatTypes,
				SeatsAvailable: *row.SectionSeatAvailable,
			}
		}

		events = append(events, event)
	}

	return events, []domain.ReviewDataModel{}, nil

	// Fetching reviews
	// reviewsQuery := fmt.Sprintf(`
	// 	SELECT
	// 		reviews.id,
	// 		reviews.rating,
	// 		reviews.created_on,
	// 		reviews.author_name,
	// 		reviews.user_id
	// 		reviews.body
	// 		reviews.image_file_path
	// 	FROM
	// 		reviews
	// 	WHERE
	// 		reviews.collection_id = '%v';
	// `, collectionId)

}
