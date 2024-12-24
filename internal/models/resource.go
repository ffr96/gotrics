package models

import "database/sql"

type ResourceType string

const (
	Banner RecordType = "banner"
	Logo   RecordType = "logo"
)

type Resource struct {
	ID       int          `json:"id"`
	RecordId int          `json:"record_id"`
	EDID     int          `json:"edid"`
	BRID     int          `json:"brid"`
	Type     ResourceType `json:"type"`
}

type ResourceModel struct {
	DB *sql.DB
}
