package models

import "database/sql"

type BrandCategory struct {
	ID   int `json:"id"`
	EDID int `json:"edid"`
	BRID int `json:"brid"`
	PCID int `json:"pcid"`
}

type BrandCategoryModel struct {
	DB *sql.DB
}

func (m *BrandCategoryModel) CreateTable() (bool, error) {

	smtm := `CREATE TABLE IF NOT EXISTS brand_category (
		id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		edid INT NOT NULL,
		brid INT NOT NULL,
		pcid INT NOT NULL
	)`

	_, err := m.DB.Exec(smtm)
	if err != nil {
		return false, err
	}

	return true, nil
}
