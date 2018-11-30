package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User))
}

// User 用户对象
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
	IsFriend              bool       `json:"isFriend" orm:"column(is_friend)"`
	IsSubscribed          bool       `json:"isSubscribed" orm:"column(is_subscribed)"`
	SubscriptionCount     int64      `json:"subscriptionCount" orm:"column(subscription_count)"`
}

// TableName 对应数据库的表名
func (u *User) TableName() string {
	return "users"
}

// Insert 插入当前用户
func (u *User) Insert() (int64, error) {
	sql := `
INSERT INTO users (yunxin_id, phone_number, name, id_card_no, id_card_avatar_url, certification_time, avatar_url, first_upload_avatar_time, create_time, update_time)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	ret, err := orm.NewOrm().Raw(sql, u.YunxinID, u.PhoneNumber, &u.Name, &u.IDCardNo, &u.IDCardAvatarURL, &u.CertificationTime, &u.AvatarURL, &u.FirstUploadAvatarTime, u.CreateTime, u.UpdateTime).Exec()
	if err != nil {
		return 0, err
	}
	return ret.LastInsertId()
}

// GetUserByPhoneNumber 通过手机号获取用户信息
func GetUserByPhoneNumber(phoneNumber string) (*User, error) {
	var user User
	sql := "SELECT *, 0 is_friend, 0 is_subscribed, 0 subscription_count FROM users WHERE phone_number = ?"
	err := orm.NewOrm().Raw(sql, phoneNumber).QueryRow(&user)
	if err == orm.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID 通过 查看人用户ID 获取某用户信息，带有两人关系数据
func GetUserByID(id int64, queryUserID int64) (*User, error) {
	var user User
	sql := `
SELECT *, 
	(SELECT id FROM friends WHERE user_id = ? AND friend_id = u.id) IS NOT NULL is_friend, 
	(SELECT id FROM subscription WHERE user_id = ? AND subscription_user_id = u.id) IS NOT NULL is_subscribed, 
	(SELECT COUNT(id) FROM subscription WHERE subscription_user_id = u.id) subscription_count 
FROM users u 
WHERE id = ?
	`
	err := orm.NewOrm().Raw(sql, queryUserID, queryUserID, id).QueryRow(&user)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}
