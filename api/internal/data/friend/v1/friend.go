package friend

import (
	"context"
	"github.com/yonymo/simplechat/pkg/common"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

const (
	Relation_Not_Add uint8 = iota
	Relation_Normal
	Relation_Delete
)

const (
	Black_List uint8 = iota + 1
	Not_Black_List
)

type FriendDO struct {
	BaseModel
	OwnerID   uint   `json:"owner_id" gorm:"index:idx_owner;type:uint"`
	FriendID  uint   `json:"friend_id" gorm:"type:uint"`
	Remark    string `json:"remark" gorm:"type:varchar(100); comment:'备注'"`
	Relation  uint8  `json:"relation" gorm:"type:tinyint;default:1; comment:'1正常 2删除 0未添加'"`
	Black     uint8  `json:"black" gorm:"type:tinyint;default:1; comment:'1正常 2拉黑'"`
	BlackSeq  string `json:"black_seq" gorm:"type:varchar(255)"`
	FriendSeq int64  `json:"friend_seq" gorm:"type:bigint"`
	AddSource string `json:"add_source" gorm:"type:varchar(20) comment:'来源'"`
	Extra     string `json:"extra" gorm:"type:varchar(1000)"`
}

func (u *FriendDO) TableName() string {
	return "friend"
}

type FriendDOList struct {
	Total int64       `json:"total,omitempty"`
	Items []*FriendDO `json:"items"`
}

type IFriendData interface {
	List(ctx context.Context, uid uint, opts common.ListMeta, orderby []string) (*FriendDOList, error)
	Create(ctx context.Context, friend *FriendDO) error
	Get(ctx context.Context, owner, dst uint) (*FriendDO, error)
	Update(ctx context.Context, friend *FriendDO) error
	//Delete(ctx context.Context, friend *FriendDO) error
	//Get(ctx context.Context, friend *FriendDO) (*FriendDO, error)
}
