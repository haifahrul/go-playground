package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseJSON struct {
	w        http.ResponseWriter
	HttpCode int
	Status   string      `json:"status"`
	Message  interface{} `json:"message"`
	Data     interface{} `json:"data"`
}

// type HTTPResponseRes struct {
// 	Status  string      `json:"status"`
// 	Message interface{} `json:"message"`
// 	Data    interface{} `json:"data"`
// }

//httpHandler Helper regarding Environment variables
type httpHandler struct {
}

//httpHandlerHandler Looks like for sending response in JSON format
func (*httpHandler) ResponseJSON(in ResponseJSON) (out *ResponseJSON) {
	var err error

	res, err := json.Marshal(in.Data)
	if err != nil {
		log.Printf("Error Marshal Response: %v \n", err)

		out.Status = "error"
		out.Message = http.StatusText(http.StatusInternalServerError)
	}

	in.w.Header().Set("Content-Type", "application/json")
	in.w.WriteHeader(in.HttpCode)
	in.w.Write(res)

	return
}

//HttpHandler Singleton of Env
var HttpHandler = &httpHandler{}
