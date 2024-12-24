package models

import (
	"database/sql"
	"time"
)

type Edition struct {
	ID      int       `json:"id"`
	Label   string    `json:"label"`
	Updated time.Time `json:"updated"`
}

type EditionModel struct {
	DB *sql.DB
}
