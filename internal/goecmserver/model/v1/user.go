package v1

import (
	uuid "github.com/satori/go.uuid"
	"go-ecm/internal/goecmserver/util/auth"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"gorm.io/gorm"
)

type User struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	UUID              uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	UserName          string    `json:"username" gorm:"comment:用户登录名"`
	Password          string    `json:"-" gorm:"comment:用户登录密码"`
	NickName          string    `json:"nickname" gorm:"default:系统用户;comment:用户昵称"`
	HeaderImg         string    `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	AuthorityId       string    `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
}

type UserList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*User `json:"items"`
}

func (u *User) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)

	return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Save(u).Error
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	//u.ExtendShadow = u.Extend.String()

	return err
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	//if err := json.Unmarshal([]byte(u.ExtendShadow), &u.Extend); err != nil {
	//	return err
	//}

	return nil
}
