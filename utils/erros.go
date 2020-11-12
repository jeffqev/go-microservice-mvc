package utils

//MicroservicError .
type MicroservicError struct {
	Menssage   string `json:"message"`
	StatusCode int    `json:"status"`
	Code       string `json:"code"`
}
