package entities

// Profile dos usuarios do sistema
type Profile struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Auth autentica o acesso do perfil
func (p *Profile) Auth() (err error) {
	return
}
