package models

import "github.com/google/uuid"

// Version represents a version of a service
type Version struct {
	ID          string `json:"id" gorm:"primaryKey"`
	ServiceID   string `json:"service_id" gorm:"index"`
	Version     string `json:"version"`
	ReleaseDate string `json:"release_date"`
}

// BeforeCreate generates a UUID for the version ID
func (v *Version) BeforeCreate() error {
	v.ID = uuid.New().String()
	return nil
}
