package controller

type ErrorResp struct {
	Description string         `json:"description"`
	Errors      []ErrorDetails `json:"errors"`
}

type ErrorDetails struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code"`
}
