package response

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ErrorResponse struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
}

type RegisterResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type LoginResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type CurrentResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}
