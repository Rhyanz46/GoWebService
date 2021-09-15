package admin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"main/utils"
	"net/http"
	"reflect"
)

type LoginData struct {
	Username 	string `json:"username"`
	Password 	string `json:"password"`
}

func (data *LoginData) Validation(body io.ReadCloser) (int, error) {
	var dataFromClient map[string]interface{}
	if body == http.NoBody{
		return http.StatusBadRequest, errors.New("body json di perlukan")
	}
	err := json.NewDecoder(body).Decode(&dataFromClient)
	if err != nil {
		return http.StatusBadRequest, errors.New("format json tidak benar")
	}

	username, err := utils.RequestValidator("username", dataFromClient, utils.RequestDataValidator{
		Type: reflect.String, Max: 10, Min: 5, Null: false,
	})
	if err != nil{
		return http.StatusBadRequest, err
	}

	password, err := utils.RequestValidator("password", dataFromClient, utils.RequestDataValidator{
		Type: reflect.Int, Max: 10, Min: 5, Null: true,
	})
	if err != nil{
		return http.StatusBadRequest, err
	}

	fmt.Println(username)
	fmt.Println(password)

	return http.StatusOK, nil
}

type RegisterData struct {
	Email 		string `json:"email"`
	FullName 	string `json:"fullname"`
	Username 	string `json:"username"`
	Password 	string `json:"password"`
}

