package user

type EmailRegisterRequest struct {
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

type EmailLoginRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UpdateAvatarRequest struct {
	Avatar []byte `json:"avatar,omitempty"`
}

type UpdateProfileRequest struct {
	Name           *string `json:"name,omitempty"`
	UserUniqueName *string `json:"userUniqueName,omitempty"`
}

type ResetUserPassword struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type User struct {
	UserID         int64   `json:"userID,string,required"`
	Name           string  `json:"name"`
	UserUniqueName string  `json:"user_unique_name"`
	Email          string  `json:"email"`
	AvatarURL      string  `json:"avatarURL"`
	ScreenName     *string `json:"screen_name"`
	UserCreateTime int64   `json:"userCreateTime"`
}
