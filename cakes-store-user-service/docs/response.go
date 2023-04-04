package docs

type ErrorResponse struct {
	Err error `json:"error"`
}
type SuccessSignIn struct {
	Token string `json:"token"`
}
type SuccessSignUp struct {
	Message string `json:"message"`
}
