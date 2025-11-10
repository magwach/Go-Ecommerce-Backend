package dto

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignUp struct {
	UserLogin
	Phone string `json:"phone"`
}

type UserVerifyCode struct {
	Code string `json:"token"`
}

type BecomeASeller struct {
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	PhoneNumber       string `json:"phone_number"`
	BankAccountNumber string `json:"bank_account_number"`
	SwiftCode         string `json:"swift_code"`
	PaymentType       string `json:"payment_type"`
}
