package db

func Insert(value string, tableName string) (string, error) {
	return "NewID", nil
}

func InsertMany(values []string, tableName string) ([]string, error) {
	return []string{"NewIDs"}, nil
}

func Query(filter string, tableName string) (string, error) {
	return filter, nil
}

func Update(value string, newValue string, tableName string) (string, error) {
	return newValue, nil
}

func UpdateMany(values []string, newValues []string, tableName string) ([]string, error) {
	return newValues, nil
}

func Delete(value string, tableName string) (bool, error) {
	return true, nil
}