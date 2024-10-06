package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

const (
	AdminUsersTable = "admin_users"
)

type TableAdminUsersModelType string

const (
	Basic  TableAdminUsersModelType = "basic"
	Senior                          = "senior"
	Nil                             = "nil"
)

func ParseTableUsersAdminModelType(s string) (TableAdminUsersModelType, error) {
	switch strings.ToLower(s) {
	case "basic":
		return Basic, nil
	case "senior":
		return Senior, nil
	case "nil":
		return Nil, nil
	default:
		return "", fmt.Errorf("invalid userType: %s", s)
	}
}

type TableAdminUsersModel struct {
	Id               uuid.UUID                `gorm:"type:uuid;primary_key;"`
	CurrentAdminType TableAdminUsersModelType `gorm:"column:admin_;not null"`
	UserType         TableAdminUsersModelType `gorm:"column:user_type;not null"`
	Email            string                   `gorm:"type:varchar(255);unique;not null"`
	Username         string                   `gorm:"type:varchar(50);unique;not null"`
	Password         string                   `gorm:"type:varchar(255);not null"`
	FirstName        string                   `gorm:"type:varchar(50);not null"`
	LastName         string                   `gorm:"type:varchar(50);not null"`
	CreatedOn        time.Time                `gorm:"column:created_on;not null"`
	UpdatedOn        time.Time                `gorm:"column:updated_on"`
}
