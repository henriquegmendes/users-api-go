package response

type UsersListResponse struct {
	Data []UserResponse `json:"data"`
	Page Page           `json:"page"`
}
