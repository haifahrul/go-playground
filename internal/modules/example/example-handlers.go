package example

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/haifahrul/go-playground/internal/helper"
)

type exampleHandler struct {
}

func (h *exampleHandler) Create(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"id":      1,
		"name":    "fahrul",
		"address": "jakarta",
		"message": "This is an example of HTTP Request with method GET and response JSON",
	}

	helper.HttpHandler.ResponseJSON(w, &helper.ResponseJSON{
		HTTPCode: 200,
		Status:   "OK",
		Message:  "Data has been create",
		Data:     data,
	})
}

func (h *exampleHandler) ReadAll(w http.ResponseWriter, r *http.Request) {
	data := make([]map[string]interface{}, 0)
	count := 10

	for i := 0; i < count; i++ {
		data = append(data, map[string]interface{}{
			"id":      i,
			"name":    "fahrul",
			"address": "jakarta",
			"message": "This is an example of HTTP Request with method GET and response JSON",
		})
	}

	helper.HttpHandler.ResponseJSON(w, &helper.ResponseJSON{
		HTTPCode: 200,
		Status:   "OK",
		Message:  "Data founded",
		Data:     data,
	})
}

func (h *exampleHandler) ReadOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if id == "" {
		helper.HttpHandler.ResponseJSON(w, &helper.ResponseJSON{
			HTTPCode: http.StatusBadRequest,
			Status:   "Error",
			Message:  "Param id is required",
			Data:     nil,
		})
		return
	}

	data := map[string]interface{}{
		"id":      1,
		"name":    "fahrul",
		"address": "jakarta",
		"message": "This is an example of HTTP Request with method GET and response JSON",
	}

	helper.HttpHandler.ResponseJSON(w, &helper.ResponseJSON{
		HTTPCode: 200,
		Status:   "OK",
		Message:  "Data founded",
		Data:     data,
	})
}

func (h *exampleHandler) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if id == "" {
		helper.HttpHandler.ResponseJSON(w, &helper.ResponseJSON{
			HTTPCode: http.StatusBadRequest,
			Status:   "Error",
			Message:  "Param id is required",
			Data:     nil,
		})
		return
	}

	data := map[string]interface{}{
		"id":      1,
		"name":    "fahrul rozi",
		"address": "Bogor",
		"message": "This is an example of HTTP Request with method GET and response JSON",
	}

	helper.HttpHandler.ResponseJSON(w, &helper.ResponseJSON{
		HTTPCode: 200,
		Status:   "OK",
		Message:  "Data has been update",
		Data:     data,
	})
}

func (h *exampleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if id == "" {
		helper.HttpHandler.ResponseJSON(w, &helper.ResponseJSON{
			HTTPCode: http.StatusBadRequest,
			Status:   "Error",
			Message:  "Param id is required",
			Data:     nil,
		})
		return
	}

	data := map[string]interface{}{
		"id":      1,
		"name":    "fahrul",
		"address": "jakarta",
		"message": "This is an example of HTTP Request with method GET and response JSON",
	}

	helper.HttpHandler.ResponseJSON(w, &helper.ResponseJSON{
		HTTPCode: 200,
		Status:   "OK",
		Message:  "Data has been delete",
		Data:     data,
	})
}

func (h *exampleHandler) Task1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	var getName = func(in string) string {
		a := "Your name is: "
		return fmt.Sprintf(a + in)
	}

	helper.HttpHandler.ResponseJSON(w, &helper.ResponseJSON{
		HTTPCode: 200,
		Status:   "OK",
		Message:  "Data task 1",
		Data:     getName(name),
	})
}

func (h *exampleHandler) Task2(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	var p *string

	var getName = func(in *string) {
		temp := fmt.Sprintf("Your name is: %v", *in)
		p = &temp
	}

	getName(&name)

	helper.HttpHandler.ResponseJSON(w, &helper.ResponseJSON{
		HTTPCode: 200,
		Status:   "OK",
		Message:  "Data task 1",
		Data:     p,
	})
}

type DataJSON struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type DataXML struct {
	ID    string `xml:"id"`
	Title string `xml:"title"`
}

func test3(in interface{}) {

}

// ExampleHandler variable singleton
var ExampleHandler = &exampleHandler{}
