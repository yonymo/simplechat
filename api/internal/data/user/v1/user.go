package user

import (
	"context"
	"github.com/yonymo/simplechat/pkg/common"
	"gorm.io/gorm"
	"time"
)

const (
	Add_No_Verify uint8 = iota + 1
	Add_Verify
)

type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserDO struct {
	BaseModel
	Mobile        string `json:"mobile" gorm:"uniqueIndex;type:varchar(11)"`
	Passwd        string `json:"passwd" gorm:"type:varchar(100);not null"`
	Nickname      string `json:"nickname" gorm:"type:varchar(50)"`
	Avatar        string `json:"avatar" gorm:"type:varchar(150)"`
	Gender        string `json:"gender" gorm:"type:varchar(6);default:male;comment:'male or female'"`
	Online        int    `json:"online" gorm:"type:tinyint"`
	Token         string `json:"token" gorm:"type:varchar(40)"`
	Memo          string `json:"memo" gorm:"type:varchar(140)"`
	FriendAddType uint8  `json:"friend_add_type" gorm:"type:tinyint; default:1; comment:'1需要验证，2无需验证'"`
}

func (u *UserDO) TableName() string {
	return "user"
}

type UserDOList struct {
	Total int64     `json:"total,omitempty"`
	Items []*UserDO `json:"items"`
}

type IUserData interface {
	List(ctx context.Context, opts common.ListMeta, orderby []string) (*UserDOList, error)
	Create(ctx context.Context, user *UserDO) error
	Update(ctx context.Context, user *UserDO) error
	GetByMobile(ctx context.Context, mobile string) (*UserDO, error)
	GetById(ctx context.Context, id uint) (*UserDO, error)
}
