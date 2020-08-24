package models

type Response struct {
	Success bool        `json:"success"`
	Msg     interface{} `json:"msg"`
	Result  interface{} `json:"result"`
}
