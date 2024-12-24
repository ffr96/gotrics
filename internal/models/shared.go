package models

func CreateTable(smtm string, m) (bool, error) {

	_, err := m.DB.Exec(smtm)
	if err != nil {
		return false, err
	}

	return true, nil
}
