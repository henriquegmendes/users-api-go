package response

type UserResponse struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
}
