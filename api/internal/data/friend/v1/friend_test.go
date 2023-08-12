package friend

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestCreateTable(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/simple-chat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Error(err.Error())
		return
	}
	db.AutoMigrate(&FriendDO{})

}
