package models

import (
	"database/sql"
	"time"
)

type Brand struct {
	ID      int       `json:"id"`
	Label   string    `json:"label"`
	Updated time.Time `json:"updated"`
}

type BrandModel struct {
	DB *sql.DB
}
