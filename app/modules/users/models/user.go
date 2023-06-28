package models

type (
	GetUserProfile struct {
		Username   string `json:"username"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Role       string `json:"role"`
		ImageUrl   string `json:"image_url"`
		ValidUntil int16  `json:"valid_until"`
		AcceptedAt string `json:"accepted_at"`
		UpdatedAt  string `json:"updated_at"`
	}
)
