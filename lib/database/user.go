package database

var db = config.db

func GetUsers() (interface{}, error) {
	var users []model.USer
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
