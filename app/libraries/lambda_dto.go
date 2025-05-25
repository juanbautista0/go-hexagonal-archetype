package libraries

type LambdaDtoResponse struct {
	Message    string      `json:"message,"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
}
