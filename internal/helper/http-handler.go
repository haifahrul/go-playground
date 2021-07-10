package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseJSON struct {
	HTTPCode int
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
func (*httpHandler) ResponseJSON(w http.ResponseWriter, in ResponseJSON) (out *ResponseJSON) {
	var err error

	res, err := json.Marshal(in.Data)
	if err != nil {
		log.Printf("Error Marshal Response: %v \n", err)

		out.Status = "error"
		out.Message = http.StatusText(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(in.HTTPCode)
	w.Write(res)

	return
}

//HttpHandler Singleton of Env
var HttpHandler = &httpHandler{}
