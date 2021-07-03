package models

import (
	"errors"
	"strings"
	"time"
)

type Property struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Name      string   `json:"name"`
	CountryID int      `json:"countryID"`
	Country   *Country `json:"-"`
}

func (p *Property) Validate() error {
	if len(strings.TrimSpace(p.Name)) == 0 {
		return errors.New("invalid property name")
	}

	return nil
}
