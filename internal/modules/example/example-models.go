package example

import "fmt"

type exampleModels struct {
}

type DataJSON struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type DataXML struct {
	ID    string `xml:"id"`
	Title string `xml:"title"`
}

func (h *exampleHandler) getName(in string) string {
	a := "RETURN. Your name is: "
	return fmt.Sprintf(a + in)
}

func (h *exampleHandler) getNamePointer(in *string) {
	temp := fmt.Sprintf("POINTER. Your name is: %v", *in)
	*in = temp
}

func (h *exampleHandler) test3(in interface{}) interface{} {
	return in
}

// ExampleModels variable singleton
var ExampleModels = &exampleModels{}
