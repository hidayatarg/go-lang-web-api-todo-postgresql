package domain

type RegisterPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Username        string `json:"username"`
}

func (d *Domain) Register(payload RegisterPayload) (*User, error){
	userExist, _ := d.DB.UserRepo.GetByEmail(payload.Email)
	if userExist != nil {
		return nil, ErrUserWithEmailAlreadyExist
	}

	userExist, _ = d.DB.UserRepo.GetByEmail(payload.Username)
	if userExist != nil {
		return nil, ErrUserWithUsernameAlreadyExist
	}
//	hash password

}
