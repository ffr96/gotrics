package models

import (
	"database/sql"
	"time"
)

type TimeDimension struct {
	ID        int       `json:"id"`
	EDID      int       `json:"edid"`
	Timestamp time.Time `json:"timestamp"`
	Updated   time.Time `json:"updated"`
}

type TimeDimensionModel struct {
	DB *sql.DB
}
