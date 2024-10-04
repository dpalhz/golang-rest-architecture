package response

type UserProfile struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRegister struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserLogin struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	SessionID string `json:"session_id"`
}

type UserDetails struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Photo    string `json:"photo,omitempty"`
}

type UserUpdate struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Photo string `json:"photo,omitempty"`
}
