package friend

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yonymo/simplechat/pkg/common"
	gin2 "github.com/yonymo/simplechat/pkg/common/gin"
	"github.com/yonymo/simplechat/pkg/log"
)

type FriendListForm struct {
	//Type     int    `form:"type" json:"type" binding:"required,min=0,max=1"`
	OwnerID uint   `form:"owner_id" json:"owner_id" binding:"required"`
	Pn      uint32 `form:"pn" json:"pn,omitempty"`
	PSize   uint32 `form:"pSize" json:"pSize,omitempty"`
}

func (f *friendServer) List(ctx *gin.Context) {
	form := &FriendListForm{}
	if err := ctx.ShouldBind(form); err != nil {
		gin2.WriteResponse(ctx, err, nil)
		return
	}
	meta := common.ListMeta{
		Pn:    form.Pn,
		PSize: form.PSize,
	}
	if form.Pn == 0 {
		meta.Pn = 1
	}
	if form.PSize == 0 {
		meta.PSize = 10
	}

	dtoList, err := f.srv.FriendSrv().GetFriendList(ctx, form.OwnerID, meta)
	if err != nil {
		gin2.WriteResponse(ctx, err, nil)
		return
	}

	data, err := json.Marshal(dtoList.Items)
	log.Debugf("marshal error: %v\n", err)
	gin2.WriteResponse(ctx, nil, gin.H{
		"total": dtoList.Total,
		"data":  string(data),
	})
}
