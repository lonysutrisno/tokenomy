package pkg

type (
	Data struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	Response struct {
		Code    int         `json:"code"`
		Message string      `json:"message,omitempty"`
		Data    interface{} `json:"data,omitempty"`
	}
)

var DummyData string = `
	[{
		"id": 1,
		"name": "A"
	}, {
		"id": 2,
		"name": "B"
	}, {
		"id": 3,
		"name": "C"
	}]
`
