package exception

import "fmt"

const (
	HttpGetFail             = iota
	RespBodyNil             = iota
	ReadBodyError           = iota
	HttpListenError         = iota
	HandleReadError         = iota
	JSONMarshalError        = iota
	CodeVerifyThreadDestroy = iota
	ImageTransError         = iota
)

type ZhiError struct {
	Code     int
	FuncName string
	Param1   string
	Param2   string
	Param3   string
}

func ErrorNotify(err ZhiError) {
	fmt.Println(err)
}
