package api

var Endpoint = struct {
	AdminLogin          string
	AdminRegister       string
	AdminDetail         string
	ChatCostumerService string
}{
	AdminLogin:          "/admin/login",
	AdminRegister:       "/admin/register",
	AdminDetail:         "/admin/detail",
	ChatCostumerService: "/chat/customer-service",
}
