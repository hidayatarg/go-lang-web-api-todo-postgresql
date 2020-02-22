package domain

type RegisterPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Username        string `json:"username"`
}

func (d *Domain) Register(payload RegisterPayload) (*User, error){
	
}
