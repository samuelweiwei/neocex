package user

import "github.com/google/uuid"

type UserAccount struct {
	Uuid                 uuid.UUID `json:"uuid" gorm:"primary_key" gorm:"type:uuid;default:uuid_generate_v4()"`
	Username             string    `json:"username" gorm:"comment:Username"`
	Password             string    `json:"password" gorm:"comment:Password"`
	PaymentPassword      string    `json:"payment_password" gorm:"comment:PaymentPassword"`
	Nickname             string    `json:"nickname" gorm:"comment:Nickname"`
	HeaderImg            string    `json:"header_img" gorm:"comment:HeaderImg"`
	Phone                string    `json:"phone" gorm:"comment:Phone"`
	Email                string    `json:"email" gorm:"comment:Email"`
	ParentId             uint      `json:"parent_id" gorm:"comment:ParentID"`
	ParentUserId         uint      `json:"parent_user_id" gorm:"comment:ParentUserID"`
	GrandParentId        uint      `json:"grand_parent_id" gorm:"comment:GrandParentID"`
	GrandParentUserId    uint      `json:"grand_parent_user_id" gorm:"comment:GrandParentUserID"`
	CountryId            uint      `json:"country_id" gorm:"comment:CountryID"`
	IdType               uint      `json:"id_type" gorm:"comment:IDType"`
	IdImages             string    `json:"id_images" gorm:"comment:IDImages"`
	AuthenticationStatus *int      `json:"authentication_status" gorm:"comment:AuthenticationStatus"`
	RealName             string    `json:"real_name" gorm:"comment:RealName"`
	IdNumber             string    `json:"id_number" gorm:"comment:IDNumber"`
	LastLoginIp          string    `json:"last_login_ip" gorm:"comment:LastLoginIP"`
	RootUserId           uint      `json:"root_user_id" gorm:"comment:RootUserID"`
	UserType             uint      `json:"user_type" gorm:"comment:UserType"`
	LastLoginTime        int64     `json:"last_login_time" gorm:"comment:LastLoginTime"`
	InviteCode           string    `json:"invite_code" gorm:"comment:InviteCode"`
	CreateAtInt          int64     `json:"create_at_int" gorm:"comment:CreateAtInt"`
}
