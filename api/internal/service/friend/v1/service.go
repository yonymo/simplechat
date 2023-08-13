package friend

import (
	"context"
	df "github.com/yonymo/simplechat/api/internal/data/friend/v1"
	"github.com/yonymo/simplechat/pkg/common"
)

type IFriendSrv interface {
	AddFriend(ctx context.Context, dto *FriendDTO) error
	GetFriendList(ctx context.Context, owner uint, meta common.ListMeta) (*FriendDTOList, error)
	GetFriend(ctx context.Context, owner, dst uint) (*FriendDTO, error)
	AddFriendReq(ctx context.Context, dto *FriendReqDTO) error
}

type friendService struct {
	friendData df.IFriendData
}

var _ IFriendSrv = &friendService{}

func NewFriendSrv(ifd df.IFriendData) IFriendSrv {
	return &friendService{friendData: ifd}
}
