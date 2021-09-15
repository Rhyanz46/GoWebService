package admin

import (
	"gorm.io/gorm"
	"main/utils"
	"net/http"
)

type Admin struct {
	DB *gorm.DB
}

func Routes(config *utils.RouterConfig) *Admin {
	return &Admin{
		DB: config.PrimaryDB,
	}
}

func (user *Admin)HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	var err error
	var status int

	if r.Method == "POST"{

		// data validation
		data := LoginData{}
		status, err = data.Validation(r.Body)
		if err != nil{
			utils.ResponseJson(w, status, utils.DataResponse{Message: err.Error()})
			return
		}


		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (user *Admin)HandleUserRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{

	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (user *Admin)HandleUserDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{

	}else if r.Method == "PUT"{

	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}