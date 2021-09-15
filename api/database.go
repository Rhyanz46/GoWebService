package api

import (
	"fmt"
	"main/database/primary"
	"main/settings"
)

func (server *Server) initializePrimaryDB() {
	server.PrimaryDB = settings.DataSettings.PrimaryDB.CreateConnection()

	//database migration
	err := server.PrimaryDB.AutoMigrate(
		&primary.User{},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
}