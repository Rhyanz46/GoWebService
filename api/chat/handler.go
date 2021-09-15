package chat

import (
	"gorm.io/gorm"
	"main/utils"
	"net/http"
)

type Chat struct {
	DB *gorm.DB
	Hub Hub
}

func Routes(config *utils.RouterConfig) *Chat {
	return &Chat{
		DB: config.PrimaryDB,
	}
}

func (user *Chat)HandleChatCustomerService(w http.ResponseWriter, r *http.Request) {
	user.Hub.ServeWs(w, r)
}