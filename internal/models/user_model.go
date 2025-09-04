package models

type LoginRequest struct {
	LoginOrEmail string `json:"login_or_email"`
	Password     string `json:"password"`
}

type RequestRegistration struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Login    string `json:"login"`
}

type RequestUpdateProfile struct {
	NikName string `json:"nik name"`
	Login   string `json:"login"`
	Id      string `json:"id"`
	Email   string `json:"email"`
}

type RequestUpdatePassword struct {
	Id          string `json:"id"`
	NowPassword string `json:"now_password"`
	NewPassword string `json:"new_password"`
}

type CreateRequest struct {
	SumSend     uint64 `json:"sum_send"`
	SumGet      uint64 `json:"sum_get"`
	Id          string `json:"id"`
	CryptWallet string `json:"crypt_wallet"`
	NumberPhone string `json:"number_phone"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type SuccessLoginResponses struct {
	Code  string `json:"code"`
	Token string `json:"token"`
	Id    string `json:"id"`
}
