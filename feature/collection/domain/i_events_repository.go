package domain

import (
	"context"

	"github.com/gofrs/uuid"
)

const (
	EventsTable   = "events"
	ReviewsTable  = "reviews"
	SectionsTable = "sections"
)

type EventDataModel struct {
	Id                   uuid.UUID        `json:"id"`
	IsSingleSectionEvent bool             `json:"isSingleSectionEvent"`
	IsIncrementTicket    bool             `json:"isIncrementTicket"`
	SectionModel         SectionDataModel `json:"sectionModel"`
	Name                 string           `json:"name"`
	Performer            string           `json:"performer"`
	EventFrom            uint64           `json:"eventFrom"`
	EventTo              uint64           `json:"eventTo"`
	Venue                string           `json:"venue"`
	Location             string           `json:"location"`
	ImagePath            string           `json:"imagePath"`
	EventTax             float64          `json:"eventTax"`             // Decimal(5,4)
	PlatformTax          float64          `json:"platformTax"`          // Decimal(5,4)
	PlatformPercentage   float64          `json:"platformPercentage"`   // Decimal(5,4)
	PlatformFixed        float64          `json:"platformFixed"`        // Decimal(10,2)
	ProcessingPercentage float64          `json:"processingPercentage"` // Decimal(5,4)
	ProcessingFixed      float64          `json:"processingFixed"`      // Decimal(10,2)
}

// This should be a User
type ReviewDataModel struct {
	Id            uuid.UUID `json:"id"`
	Rating        int8      `json:"rating"`
	Date          uint64    `json:"date"`
	Author        string    `json:"author"`
	AuthorId      string    `json:"authorId"`
	Body          string    `json:"body"`
	ImageFilePath string    `json:"imageFilePath"`
}

type SectionDataModel struct {
	Id             uuid.UUID `json:"id"`
	SeatingMap     string    `json:"seatingMap"`
	Name           string    `json:"name"`
	LowestPrice    float32   `json:"lowestPrice"`
	HighestPrice   float32   `json:"highestPrice"`
	SeatTypes      []string  `json:"seatTypes"`
	SeatsAvailable bool      `json:"seatsAvailable"`
}

type EventsRepository interface {
	GetCollectionHeavyData(
		c context.Context,
		collectionId string,
		domainAddress string,
	) ([]EventDataModel, []ReviewDataModel, error)
}
