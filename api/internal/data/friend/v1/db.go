package friend

import (
	"context"
	"github.com/yonymo/simplechat/pkg/common"
)

type IFriendData interface {
	List(ctx context.Context, uid uint, opts common.ListMeta) (*FriendDOList, error)
	Create(ctx context.Context, friend *FriendDO) error
	Get(ctx context.Context, owner, dst uint) (*FriendDO, error)
	Update(ctx context.Context, friend *FriendDO) error
	//Delete(ctx context.Context, friend *FriendDO) error
	//Get(ctx context.Context, friend *FriendDO) (*FriendDO, error)

	CreateFriendReq(ctx context.Context, friendReq *FriendReqDO) error
	GetFriendReq(ctx context.Context, owner, dst uint) (*FriendReqDO, error)
	UpdateFriendReq(ctx context.Context, friendReq *FriendReqDO) error
}
