package helper

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func ApiResponse(message string, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Status:  status,
	}

	jsonResp := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResp
}
