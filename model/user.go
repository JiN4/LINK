package model

// Create userを作成する
func CreateUser(user User) (User, error) {
	err := db.Create(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
