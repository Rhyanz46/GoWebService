package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"main/api"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	admin             = "admin"
	user              = "user"
	unauthorized      = "unauthorized"
	unknownAuthorized = "unknownAuthorized"

	all  = "ALL"
	get  = "GET"
	post = "POST"
	put  = "PUT"
)

var LoginExpirationDuration = time.Duration(1) * time.Hour
var JwtSigningMethod = jwt.SigningMethodHS256
var JwtSignatureKey = []byte("FO00XG0P0ndas1CrOkBos55JAYUDA44e4ateByVNEU@20200218Gara2BCA")

var permissionList = map[string]map[string][]string{
	admin: {
		api.Endpoint.AdminDetail:         {all},
		api.Endpoint.ChatCostumerService: {all},
	},
	user: {
		api.Endpoint.AdminDetail:         {all},
		api.Endpoint.ChatCostumerService: {all},
	},
	unknownAuthorized: {
		api.Endpoint.AdminLogin: {get, post},
	},
	unauthorized: {
		api.Endpoint.AdminLogin: {get, post},
	},
}

type Auth struct {
	Name    string
	Phone   string
	Email   string
	Role    int
	Project string
}

func PermissionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var role string
		var status int
		var authData Auth
		bearerToken := getBearerToken(r)
		if bearerToken != "" {
			status, authData = GetTokenData(bearerToken)
			if status != http.StatusOK {
				w.WriteHeader(status)
				return
			}
			if authData.Role == 0 {
				role = admin
			} else if authData.Role == 1 {
				role = user
			} else {
				role = unknownAuthorized
			}
		} else {
			role = unauthorized
		}

		if !canAccess(r, role) {
			w.WriteHeader(http.StatusUnauthorized)
			err := json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
			}{
				Message: "Anda tidak bisa mengakses ini",
			})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetTokenData(token string) (int, Auth) {
	if token == "" {
		return http.StatusBadRequest, Auth{}
	}
	bearerToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return JwtSignatureKey, nil
	})
	if err != nil || !bearerToken.Valid {
		return http.StatusUnauthorized, Auth{}
	}
	dat := bearerToken.Claims.(jwt.MapClaims)
	iRole, _ := strconv.Atoi(fmt.Sprintf("%v", dat["Role"]))
	return http.StatusOK, Auth{
		Name:    fmt.Sprintf("%v", dat["Name"]),
		Phone:   fmt.Sprintf("%v", dat["Phone"]),
		Email:   fmt.Sprintf("%v", dat["Email"]),
		Role:    iRole,
		Project: fmt.Sprintf("%v", dat["Project"]),
	}
}

func getBearerToken(r *http.Request) string {
	var token string
	token = r.Header.Get("Authorization")
	if token != "" {
		splitToken := strings.Split(token, " ")
		if len(splitToken) == 2 {
			if strings.ToLower(splitToken[0]) != "bearer" {
				return ""
			}
			return splitToken[1]
		}
	}
	token = r.Header.Get("authorization")
	if token != "" {
		splitToken := strings.Split(token, " ")
		if len(splitToken) == 2 {
			if strings.ToLower(splitToken[0]) != "bearer" {
				return ""
			}
			return splitToken[1]
		}
	}
	return ""
}

func canAccess(r *http.Request, role string) (access bool) {
	ruote, _ := mux.CurrentRoute(r).GetPathTemplate()
	getMethod := r.Method
	permission := permissionList[role]
	for url := range permission {
		if ruote == url {
			for _, method := range permission[url] {
				if method == all {
					access = true
					goto returnResult
				}
				if method == getMethod {
					access = true
					goto returnResult
				}
			}
		}
	}
returnResult:
	return
}
