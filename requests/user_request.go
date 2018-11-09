package requests

type UserRequest struct {
	GetList struct {
		Page    int `json:"page" form:"page"`
		PerPage int `json:"per_page" form:"per_page"`
	}
}
