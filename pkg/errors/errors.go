package rest

type HttpErrorResponse struct {
	Error       error  `json:"error"`
	Explanation string `json:"explanation,omitempty"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
