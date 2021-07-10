package example

import "github.com/haifahrul/go-playground/internal/router"

var routes = []router.Route{
	{Name: "ExampleCreate", Methods: []string{"POST"}, Pattern: "/example", HandlerFunc: ExampleHandler.Create},
	{Name: "ExampleReadAll", Methods: []string{"GET"}, Pattern: "/example", HandlerFunc: ExampleHandler.ReadAll},
	{Name: "ExampleReadOne", Methods: []string{"GET"}, Pattern: "/example/{id}", HandlerFunc: ExampleHandler.ReadOne},
	{Name: "ExampleUpdate", Methods: []string{"PUT"}, Pattern: "/example/{id}", HandlerFunc: ExampleHandler.Update},
	{Name: "ExampleDelete", Methods: []string{"DELETE"}, Pattern: "/example/{id}", HandlerFunc: ExampleHandler.Delete},
}

func init() {
	router.RouteApply(routes)
}
