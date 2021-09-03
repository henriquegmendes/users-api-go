package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"henrique.mendes/users-api/database"
	"henrique.mendes/users-api/dtos/request"
	"henrique.mendes/users-api/mappers"
	"henrique.mendes/users-api/models"
)

func Paginate(page int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

func Create(user models.User) (models.User, error) {
	err := database.DB.Create(&user).Error

	if err != nil {
		return models.User{}, errors.New(err.Error())
	}

	return user, nil
}

func FindByEmail(email string) models.User {
	var user models.User
	database.DB.Where("email = ?", email).First(&user)

	return user
}

func FindById(userId uint) models.User {
	var user models.User
	database.DB.Where("id = ?", userId).First(&user)

	return user
}

func FindByNamePaginated(name string, page int, limit int) ([]models.User, int) {
	var user []models.User
	var total int64
	sql := fmt.Sprintf("'%s' = '' or name like '%%%s%%'", name, name)

	database.DB.Scopes(Paginate(page, limit)).Where(sql).Find(&user)
	database.DB.Model(&user).Where(sql).Count(&total)

	return user, int(total)
}

func UpdateUserById(userId uint64, data request.UserUpdateRequest) models.User {
	user := mappers.ToUpdateUser(data)

	database.DB.Where("id = ?", userId).Updates(user).Scan(&user)

	return user
}

func DeleteUserById(userId uint64) error {
	return database.DB.Delete(&models.User{}, userId).Error
}
