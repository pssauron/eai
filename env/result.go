package context

const (
	RESULT_OK     = "0"
	RESULT_OK_MSG = "调用成功"
)

type Result struct {
	Code    string      `json xml:"code"`
	Message string      `json xml:"message"`
	Data    interface{} `json xml:"data"`
}

func NewResult(code, message string, data interface{}) *Result {
	return &Result{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
