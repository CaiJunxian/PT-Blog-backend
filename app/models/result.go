package models

const (
	SuccessCode      = "0"
	DefaultErrorCode = "1"
	UnknownErrorCode = "-1"
	SuccessMsg       = "success"
	DefaultErrorMsg  = "error"
)

type Result struct {
	Code string
	Data interface{}
	Msg  string
}

func Success(data interface{}) Result {
	r := &Result{}
	r.Code = SuccessCode
	r.Data = data
	r.Msg = SuccessMsg
	return *r
}

func Error(code, msg string) Result {
	r := &Result{}
	r.Code = code
	r.Data = nil
	r.Msg = msg
	return *r
}
