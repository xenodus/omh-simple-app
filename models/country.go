package models

import (
	"errors"
	"strings"
	"time"
)

type Country struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"unique" json:"name"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (c *Country) Validate() error {
	if len(strings.TrimSpace(c.Name)) == 0 {
		return errors.New("invalid country name")
	}

	return nil
}
