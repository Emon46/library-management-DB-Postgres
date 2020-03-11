package pkg

type TheData interface{}

type MyData struct {
	Success string  `json:"success"`
	Status  int     `json:"status"`
	Error   string  `json:"error"`
	Message string  `json:"message"`
	Data    TheData `json:"data"`
}
