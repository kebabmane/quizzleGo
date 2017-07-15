package models

import "github.com/go-ozzo/ozzo-validation"

// Artist represents an artist record.
type Fact struct {
	Id         int    `json:"id" db:"id"`
	FactString string `json:"factString" db:"factString"`
}

// Validate validates the Artist fields.
func (m Fact) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.FactString, validation.Required, validation.Length(0, 1000)),
	)
}
