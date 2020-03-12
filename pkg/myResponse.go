package pkg

type TheData interface{}

type MyData struct {
	Success string  `json:"success"`
	Status  int     `json:"status"`
	Error   error   `json:"error"`
	Message string  `json:"message"`
	Data    TheData `json:"data"`
}
