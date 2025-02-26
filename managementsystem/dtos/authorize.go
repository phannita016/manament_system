package dtos

type (
	LoginRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	RefreshTokenRequest struct {
		Token string `json:"token" validate:"required"`
	}

	RefreshTokenResponse struct {
		AccessToken string `json:"access_token"`
	}
)
