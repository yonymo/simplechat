package db

import (
	"github.com/yonymo/simplechat/api/internal/data/friend/v1"
	"gorm.io/gorm"
)

type Friend struct {
	db *gorm.DB
}

var _ friend.IFriendData = &Friend{}

func NewFriendData(db *gorm.DB) friend.IFriendData {
	return &Friend{db: db}
}
