package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/klaus01/twoline/models"
	"github.com/klaus01/twoline/until"
)

func Login(phoneNumber string) (*models.User, error) {
	user, err := models.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}
	// 注册
	user = new(models.User)
	user.YunxinID = "yx_" + phoneNumber
	user.PhoneNumber = phoneNumber
	user.CreateTime = until.Now()
	user.UpdateTime = until.Now()

	o := orm.NewOrm()
	id, err := o.Insert(user)
	if err != nil {
		return nil, err
	}
	user.ID = id
	return user, nil
}
