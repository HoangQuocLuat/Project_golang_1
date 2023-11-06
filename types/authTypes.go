package types

type RegisterRes struct {
	Result Result          `json:"result"`
	Data   UserRegisterRes `json:"data"`
}
type UserRegisterReq struct {
	Username string `json: "username"`
	Password string `json: "password"`
	Fullname string `json: "fullname"`
	Email    string `json: "email"`
}

type UserRegisterRes struct {
	Username string `json: "username"`
	Password string `json: "password"`
	Fullname string `json: "fullname"`
	Email    string `json: "email"`
}

type Result struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
