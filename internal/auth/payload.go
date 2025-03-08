package auth

type(
	//register
	RegisterRequest struct{
		Name string `json:"name" validate:"required,min=5"`
		Email string `josn:"email" validate:"required,email,min=8"`
		Password string `json:"password" validate:"required,min=10"`
	}

	RegisterResponse struct{
		Token string `json:"token"`
	}
	//login
	LoginRequest struct{	
		Email string `josn:"email" validate:"required,email,min=8"`
		Password string `json:"password" validate:"required,min=10"`
	}

	LoginResponse struct{
		Token string `json:"token"`
	}
)