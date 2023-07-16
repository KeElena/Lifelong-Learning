package errorx

// ParamsError 自定义错误代码一般要求微服务编号+错误信息代码
//101表示微服务编号，01为错误代码
var ParamsError = New(10101, "自定义错误")

// BizError 自定义错误结构体
type BizError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// New 构造方法
func New(code int, msg string) *BizError {
	return &BizError{Code: code, Msg: msg}
}

//返回错误字符串，实现error接口方法
func (e *BizError) Error() string {
	return e.Msg
}

// ErrorResponse 返回前端的响应消息结构体
type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *BizError) Data() *ErrorResponse {
	return &ErrorResponse{
		Msg:  e.Msg,
		Code: e.Code,
	}
}
