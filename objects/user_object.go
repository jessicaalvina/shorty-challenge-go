package objects

type UserObject struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	ImageProfile string `json:"image_profile"`
}
