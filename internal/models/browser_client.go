package models

import "database/sql"

type Device string

const (
	Mobile  Device = "mobile"
	Desktop Device = "desktop"
	Tablet  Device = "tablet"
)

type BrowserClient struct {
	ID              int    `json:"id"`
	BrowserClientId string `json:"browser_client_id"`
	Ip              string `json:"ip"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	Device          Device `json:"device"`
}

type BrowserClientModel struct {
	DB *sql.DB
}

// CREATE TABLE BrowserClient (
//     id INT PRIMARY KEY AUTO_INCREMENT,
//     browser_client_id VARCHAR(255) NOT NULL,
//     ip VARCHAR(45) NOT NULL, -- Supports both IPv4 and IPv6
//     width INT NOT NULL,
//     height INT NOT NULL,
//     device ENUM('mobile', 'desktop', 'tablet') NOT NULL
// );
