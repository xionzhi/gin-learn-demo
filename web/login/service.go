package login

import (
	models "apps/models"
	utils "apps/utils"
)

var UserFields = []string{"id", "password", "account", "email"}

func QueryOneUsersByAccount(account string) (*models.User, error) {
	userOne := &models.User{}
	err := utils.DBSession.Select(UserFields).
		Where("account = ?", account).
		First(&userOne).Error

	if err != nil {
		return nil, err
	} else {
		return userOne, nil
	}
}
