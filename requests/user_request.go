package requests

type UserRequest struct {
	GetList struct {
		Page    int `json:"page" form:"page"`
		PerPage int `json:"per_page" form:"per_page"`
	}
	UpdateUser struct {
		Name  int `json:"name" form:"name"`
		Email int `json:"email" form:"email"`
	}
}
