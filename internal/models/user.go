package models

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}
