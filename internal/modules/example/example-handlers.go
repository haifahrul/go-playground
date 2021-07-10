package example

import (
	"net/http"

	"github.com/gorilla/mux"
)

type exampleHandler struct {
}

func (h *exampleHandler) Get(w http.ResponseWriter, r *http.Request) {
	// var str string
	// if r.Method == "POST" {
	// 	var bodyBytes []byte
	// 	bodyBytes, _ = ioutil.ReadAll(r.Body)
	// 	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// 	str = string(bodyBytes)
	// }
	// str = "Hasilnya: " + str
	// w.Write([]byte(str))

	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		// helper.HttpHandler.ResponseJSON(h)

		// w, http.StatusBadRequest, "error", "Param id is required", nil
		return
	}
}

// ExampleHandler variable singleton
var ExampleHandler = &exampleHandler{}
