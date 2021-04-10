package dto

// DTO 数据传输曾

type MessageReslut struct {
	Result  interface{} `json:"result"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}
