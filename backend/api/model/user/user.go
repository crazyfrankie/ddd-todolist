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
	UserUniqueName *string `json:"user_unique_name,omitempty"`
}

type ResetUserPassword struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type User struct {
	UserIDStr      int64   `json:"user_id_str,string,required"`
	Name           string  `json:"name"`
	UserUniqueName string  `json:"user_unique_name"`
	Email          string  `json:"email"`
	AvatarURL      string  `json:"avatar_url"`
	ScreenName     *string `json:"screen_name"`
	UserCreateTime int64   `json:"user_create_time"`
}
