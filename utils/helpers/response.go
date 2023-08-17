package helpers

type Response struct {
	Code    string `json:"code,omitempty"`
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

/*
code: error code from code_feature+code_layer+http_response_code.
message: message response.
*/
func WebResponseError(messages any, code string) map[string]any {
	response := map[string]any{
		"meta": Response{
			Code:    code,
			Message: "failed",
		},
		"messages": messages,
	}
	return response
}

/*
code: error code from code_feature+code_layer+http_response_code.
message: message response.
data: if data response is available, put here. otherwise put nil.
*/
func WebResponseSuccess(messages any, code string, data any) map[string]any {
	response := map[string]any{
		"meta": Response{
			Code:   code,
			Status: "success",
		},
		"messages": messages,
		"data":     data,
	}
	return response
}
