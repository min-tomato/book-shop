package response

const (
	ErrCodeSuccess       = 20001 // success
	ErrCodeParamInvalid  = 20003 // email is invalid
	ErrInvalidToken      = 30001 // token is invalid
	ErrCodeUserHasExists = 50001 // user has already registered
)

var msg = map[int]string{
	ErrCodeSuccess:       "Success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrInvalidToken:      "Token is invalid",
	ErrCodeUserHasExists: "User has already registered",
}
