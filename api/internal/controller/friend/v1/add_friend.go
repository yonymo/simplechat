package friend

import (
	"github.com/gin-gonic/gin"
	"github.com/yonymo/simplechat/api/internal/data/user/v1"
	"github.com/yonymo/simplechat/api/internal/service/friend/v1"
	"github.com/yonymo/simplechat/pkg/code"
	gin2 "github.com/yonymo/simplechat/pkg/common/gin"
	"github.com/yonymo/simplechat/pkg/errors"
)

const (
	ADD_BY_ID int = iota
	ADD_BY_NICKNAME
)

type AddFriendForm struct {
	//Type     int    `form:"type" json:"type" binding:"required,min=0,max=1"`
	*friend.FriendDTO
}

func (f *friendServer) AddFriend(ctx *gin.Context) {
	addForm := &AddFriendForm{}
	if err := ctx.ShouldBind(addForm); err != nil {
		gin2.WriteResponse(ctx, errors.WithCode(code.ErrParam, ""), nil)
		return
	}

	_, err := f.srv.UserSrv().Get(ctx, addForm.OwnerID)
	if err != nil {
		gin2.WriteResponse(ctx, err, nil)
		return
	}
	friendDto, err := f.srv.UserSrv().Get(ctx, addForm.FriendID)
	if err != nil {
		gin2.WriteResponse(ctx, err, nil)
		return
	}
	if friendDto.FriendAddType == user.Add_No_Verify {
		err := f.srv.FriendSrv().AddFriend(ctx, addForm.FriendDTO)
		if err != nil {
			gin2.WriteResponse(ctx, err, nil)
			return
		}
	} else {
		// TODO: 验证好友申请
	}

	gin2.WriteResponse(ctx, nil, gin.H{})
}
