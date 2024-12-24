package models

import (
	"database/sql"
	"time"
)

type RecordType string

const (
	BannerImpression RecordType = "banner_impression"
	BannerClick      RecordType = "banner_click"
	LogoImpression   RecordType = "logo_impression"
	LogoClick        RecordType = "logo_click"
	StartSession     RecordType = "start_session"
	NewClient        RecordType = "new_client"
)

type BrowserRecord struct {
	ID         int        `json:"id"`
	TDID       int        `json:"tdid"`
	CRID       int        `json:"crid"`
	RecordType RecordType `json:"type"`
	RSID       int        `json:"rsid"`
	Qty        int        `json:"qty"`
	Updated    time.Time  `json:"updated"`
}

type BrowserRecordModel struct {
	DB *sql.DB
}
