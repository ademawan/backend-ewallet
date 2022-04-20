package auth

type LoginReqFormat struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginRespFormat struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type UserLoginResponse struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

// type UserLoginGoogle struct {
// 	Id             string `json:"id"`
// 	Email          string `json:"email"`
// 	Verified_email bool   `json:"verified_email"`
// 	Name           string `json:"name"`
// 	Given_name     string `json:"given_name"`
// 	Family_name    string `json:"family_name"`
// 	Picture        string `json:"picture"`
// 	Locale         string `json:"locale"`
// }
