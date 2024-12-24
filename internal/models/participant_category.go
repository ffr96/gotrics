package models

import (
	"database/sql"
	"time"
)

type ParticipantCategory struct {
	ID      int       `json:"id"`
	Label   string    `json:"label"`
	Updated time.Time `json:"updated"`
}

type ParticipantCategoryModel struct {
	DB *sql.DB
}
