package main

func GetHistorySpDetail(tx_hash string) (string, error) {
	row := db.Table("txes").
		Select("property_data").
		Joins("INNER JOIN smart_properties ON tx_id = create_tx_id").
		Where("tx_hash = ?", tx_hash).
		Row()

	var propertyData string
	err := row.Scan(&propertyData)
	if err != nil {
		return "", err
	}

	return propertyData, nil
}
