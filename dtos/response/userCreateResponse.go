package response

type UserCreateResponse struct {
	Id      uint   `json:"Id"`
	Name    string `json:"Name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
}
