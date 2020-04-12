package httpserver

import (
	"demo/configure"
	"demo/exception"
	"demo/orm"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

var handlerSets = map[string]interface{}{
	//静态资源
	"/index": rootHandel,
	"/":      rootHandel,
	"/css/":  cssHandel,
	//动态接口
	"/query/": queryHandel,
}

func queryHandel(writer http.ResponseWriter, request *http.Request) {
	sqlConditional := request.URL.Query().Get("sql")
	var dataLen int
	var retJSON JSONResponse
	dataLen = len(orm.SelectUserSizeByConditional(sqlConditional))
	retJSON = JSONResponse{Data: &DataResponse{Size: dataLen}}
	result, err := json.Marshal(retJSON)
	if err != nil {
		exception.ErrorNotify(exception.ZhiError{
			Code:     exception.JSONMarshalError,
			FuncName: "queryHandel",
		})
	}
	writer.Write(result)
}
func cssHandel(writer http.ResponseWriter, request *http.Request) {
	bytes, err := ioutil.ReadFile("httpserver/html" + request.URL.Path)
	if err == nil {
		writer.Header().Add("Content-Type", "text/css")
		writer.WriteHeader(200)
		writer.Write(bytes)
	} else {
		panic(exception.ZhiError{
			Code:     exception.HandleReadError,
			FuncName: "cssHandel",
		})
	}
}
func rootHandel(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("httpserver/html/index.html")
	if err == nil {
		t.Execute(writer, nil)
	} else {
		panic(exception.ZhiError{
			Code:     exception.HandleReadError,
			FuncName: "rootHandel",
		})
	}
}
func addHandler(serMux *http.ServeMux) {
	for path, v := range handlerSets {
		serMux.HandleFunc(path, v.(func(w http.ResponseWriter, r *http.Request)))
	}
}
func StartWebServer() {
	serMux := http.NewServeMux()
	addHandler(serMux)
	serConfig := http.Server{
		Addr:              configure.GetHttpServerAddressPort(),
		Handler:           serMux,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	err := serConfig.ListenAndServe()
	if err != nil {
		panic(exception.ZhiError{
			Code:     exception.HttpListenError,
			FuncName: "StartWebServer",
			Param1:   configure.GetHttpServerAddressPort(),
			Param2:   "",
			Param3:   "",
		})
	}
}
