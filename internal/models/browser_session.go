package models

import (
	"database/sql"
	"time"
)

type BrowserSession struct {
	ID               int       `json:"id"`
	TDID             int       `json:"tdid"`
	BCID             int       `json:"bcid"`
	BrowserSessionId string    `json:"browser_session_id"`
	Updated          time.Time `json:"updated"`
}

type BrowserSessionModel struct {
	DB *sql.DB
}
