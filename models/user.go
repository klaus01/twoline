package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	ID                    int64      `json:"id" orm:"column(id)"`
	YunxinID              string     `json:"yunxinID" orm:"column(yunxin_id)"`
	PhoneNumber           string     `json:"phoneNumber"`
	Name                  *string    `json:"name"`
	IDCardNo              *string    `json:"idCardNo" orm:"column(id_card_no)"`
	IDCardAvatarURL       *string    `json:"idCardAvatarURL" orm:"column(id_card_avatar_url)"`
	CertificationTime     *time.Time `json:"certificationTime"`
	AvatarURL             *string    `json:"avatarURL" orm:"column(avatar_url)"`
	FirstUploadAvatarTime *time.Time `json:"firstUploadAvatarTime"`
	CreateTime            time.Time  `json:"createTime" orm:"auto_now_add;type(datetime)"`
	UpdateTime            time.Time  `json:"updateTime" orm:"auto_now;type(datetime)"`
}

func (u *User) TableName() string {
	return "users"
}

func GetUserByPhoneNumber(phoneNumber string) (*User, error) {
	var user User
	o := orm.NewOrm()
	err := o.QueryTable(user).Filter("phone_number", phoneNumber).One(&user)
	if err == orm.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
