package helper

// Response struct is representing ..
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// Meta struct is representing ..
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

// APIResponse func is intended to create standarized response from the api to the clients
func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonresponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonresponse
}
