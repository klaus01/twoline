package services

import (
	"github.com/klaus01/twoline/models"
	"github.com/klaus01/twoline/until"
)

// Login 登录业务
func Login(phoneNumber string) (*models.User, bool, error) {
	user, err := models.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, false, err
	}
	if user != nil {
		return user, false, nil
	}
	// 注册
	user = new(models.User)
	user.YunxinID = "yx_" + phoneNumber
	user.PhoneNumber = phoneNumber
	user.CreateTime = until.Now()
	user.UpdateTime = until.Now()

	id, err := user.Insert()
	if err != nil {
		return nil, false, err
	}
	user.ID = id
	return user, true, nil
}
