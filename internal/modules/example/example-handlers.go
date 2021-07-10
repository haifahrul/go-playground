package example

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/haifahrul/go-playground/internal/helper"
)

type exampleHandler struct {
}

func (h *exampleHandler) Get(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"name":    "fahrul",
		"address": "jakarta",
		"message": "This is an example of HTTP Request with method GET and response JSON",
	}

	helper.HttpHandler.ResponseJSON(w, helper.ResponseJSON{
		HTTPCode: 500,
		Status:   "OK",
		Message:  "Success",
		Data:     data,
	})
}

func (h *exampleHandler) GeByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if id == "" {
		helper.HttpHandler.ResponseJSON(w, helper.ResponseJSON{
			HTTPCode: http.StatusBadRequest,
			Status:   "Error",
			Message:  "Param id is required",
			Data:     nil,
		})
		return
	}

	helper.HttpHandler.ResponseJSON(w, helper.ResponseJSON{
		HTTPCode: 200,
		Status:   "OK",
		Message:  "Success",
		Data:     nil,
	})
}

// ExampleHandler variable singleton
var ExampleHandler = &exampleHandler{}
