package auth

type(
	//register
	RegisterRequest struct{
		Name string `json:"name" validate:"required,min=3"`
		Email string `josn:"email" validate:"required,email,min=8"`
		Password string `json:"password" validate:"required,min=10"`
	}
	RegisterResponse struct{
		Token string `json:"token"`
	}
	RegisterCreateResponse struct{
		Msg string `json:"msg"`
		CreateUser bool `json:"CreateUser"`
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