package exception

const (
	HttpGetFail   = iota
	RespBodyNil   = iota
	ReadBodyError = iota
)

type ZhiError struct {
	Code     int
	FuncName string
	Param1   string
	Param2   string
	Param3   string
}
