package models

import (
	"time"
)

// Text -> Text struct to save Text on database
type Text struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:200" json:"title"`
	Body      string    `gorm:"size:3000" json:"body" `
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// TableName method that returns tablename of Text model
func (text *Text) TableName() string {
	return "text"
}

// ResponseMap -> response map of Text
func (text *Text) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = text.ID
	resp["title"] = text.Title
	resp["body"] = text.Body
	resp["created_at"] = text.CreatedAt
	resp["updated_at"] = text.UpdatedAt
	return resp
}
