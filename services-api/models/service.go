package models

import "github.com/google/uuid"

// Service represents a service in the catalog
type Service struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Versions    []Version `json:"-" gorm:"foreignKey:ServiceID"` // Excluded from JSON output
}

// BeforeCreate generates a UUID for the service ID
func (s *Service) BeforeCreate() error {
	s.ID = uuid.New().String()
	return nil
}
