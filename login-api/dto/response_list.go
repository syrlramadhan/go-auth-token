package dto

type ResponseList struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ResponseToken struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Token   string `json:"token"`
	Message string `json:"message"`
}
