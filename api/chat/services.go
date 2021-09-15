package chat

import (
	"fmt"
	"log"
	"net/http"
)


func webSocketCustomerService(w http.ResponseWriter, r *http.Request)  {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(ws)
}