package models

const (
	UserStateActive  = "active"
	UserStateBlocked = "blocked"
)

// User 用户
type User struct {
	BaseModel
	Email    string `gorm:"type:varchar(100);unique_index;not null;" json:"email"`   // 用户通过邮箱地址进行登陆
	UserName string `gorm:"type:varchar(50);" json:"user_name"`                      // 用户名
	Name     string `gorm:"type:varchar(30);not null" json:"name"`                   // 真实姓名
	State    string `gorm:"type:varchar(20);default:'active';not null" json:"state"` // 用户状态 active:启用，blocked:禁用, blocked:ldap禁用
	Mobile   string `gorm:"type:varchar(11)" json:"mobile"`                          // 手机号
	IsAdmin  *bool  `gorm:"default:false" json:"is_admin"`                           // 是否为管理员
	Avatar   string `gorm:"type:varchar(255);" json:"avatar"`                        // 用户头像
}

func (u *User) IsActive() bool {
	return u.State == UserStateActive
}
