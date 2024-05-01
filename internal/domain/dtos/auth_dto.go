package dtos

type UserAuthDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

type UserAuthResponseDTO struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func NewUserAuthResponseDTO(token, refreshToken string) *UserAuthResponseDTO {
	return &UserAuthResponseDTO{
		Token:        token,
		RefreshToken: refreshToken,
	}
}

type UserAuthRefreshDTO struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
