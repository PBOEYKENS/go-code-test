package domain

import (
	"time"

	"github.com/gofrs/uuid"
)

type Base struct {
	Id            uuid.UUID            `gorm:"type:uuid;primary_key;"`
	CreatedOn     time.Time            `gorm:"column:created_on;not null"`
	UpdatedOn     time.Time            `gorm:"column:updated_on"`
	CreatedBy     uuid.UUID            `gorm:"column:created_by; not null"`
	UpdatedBy     uuid.UUID            `gorm:"column:updated_by"`
	FkCreatedById TableAdminUsersModel `gorm:"constraint:fk_created_by_id;foreignKey:CreatedBy;references:Id"`
	FkUpdatedById TableAdminUsersModel `gorm:"constraint:fk_updated_by_id;foreignKey:UpdatedBy;references:Id"`
}

func CreateBase(adminUUID *uuid.UUID) (Base, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return Base{}, err
	}

	currentTime := time.Now().UTC()

	return Base{
		Id:        id,
		CreatedOn: currentTime,
		UpdatedOn: currentTime,
		CreatedBy: *adminUUID,
		UpdatedBy: *adminUUID,
	}, nil
}

func CreateUpdateBase(adminUUID *uuid.UUID) (Base, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return Base{}, err
	}

	return Base{
		Id:        id,
		UpdatedOn: time.Now().UTC(),
		UpdatedBy: *adminUUID,
	}, nil
}

func CreateBaseSortableId(adminUUID uuid.UUID) (Base, error) {
	id, err := uuid.NewV6()
	if err != nil {
		return Base{}, err
	}

	return Base{
		Id:        id,
		CreatedOn: time.Now().UTC(),
		CreatedBy: adminUUID,
	}, nil
}
