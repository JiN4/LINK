package model

// Create browsedHistoryを作成する
func CreateBrowsedHistory(browsedHistory BrowsedHistory) (BrowsedHistory, error) {
	err := db.Create(&browsedHistory).Error
	if err != nil {
		return BrowsedHistory{}, err
	}
	return browsedHistory, nil
}
