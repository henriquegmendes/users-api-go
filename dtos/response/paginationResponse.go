package response

type Page struct {
	Page         int `json:"page"`
	TotalPerPage int `json:"page_total"`
	TotalResults int `json:"total_results"`
	LastPage     int `json:"last_page"`
}
