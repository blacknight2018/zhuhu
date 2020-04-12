package httpserver

type DataResponse struct {
	Size int `json:"size"`
}
type JSONResponse struct {
	Data *DataResponse `json:"data"`
}
