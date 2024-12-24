package models

import "database/sql"

type Container struct {
	ID   int    `json:"id"`
	EDID int    `json:"edid"`
	Name string `json:"name"`
}

type ContainerModel struct {
	DB *sql.DB
}
