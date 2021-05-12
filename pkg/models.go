package pkg

type (
	Data struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	Response struct {
		Code    int         `json:"code"`
		Message string      `json:"message,omitempty"`
		Data    interface{} `json:"data"`
	}
)
