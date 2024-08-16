package models

const InvalidCredentialsErr = "user not registered or bad credentials"
const EmailAlreadyInUseErr = "this email is already registered"
const DbErr = "generic db err"
const ObjectNotFoundErr = "the object with the requested id is not found"
const ValidationErr = "body validation"
const TokenGenerationErr = "failure in generating JWT Token"
const DateWrongFormatErr = "date has a wrong format. Expected YYYY-mm-DD"
const MissingTokenErr = "the jwt token is missing"
const TokenNotValidErr = "the jwt token is not valid"

type CoworkingError struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (c CoworkingError) Error() string {
	return c.Message
}
