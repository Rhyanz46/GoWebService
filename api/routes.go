package api

import (
	"main/api/admin"
	"main/api/chat"
	"main/utils"
)

func (server *Server) initializeRoutes() {
	route := server.Router.HandleFunc

	// admin router
	userRoute := admin.Routes(&utils.RouterConfig{PrimaryDB: server.PrimaryDB})
	route("/admin/login", userRoute.HandleUserLogin)
	route("/admin/register", userRoute.HandleUserRegister)
	route("/admin/detail", userRoute.HandleUserDetail)

	// chat router
	chatRoute := chat.Routes(&utils.RouterConfig{PrimaryDB: server.PrimaryDB})
	chatRoute.Hub.Run()
	route("/chat/customer-service", chatRoute.HandleChatCustomerService)

}