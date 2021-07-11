package example

import "github.com/haifahrul/go-playground/internal/router"

var routes = []router.Route{
	{Name: "ExampleCreate", Methods: []string{"POST"}, Pattern: "/example", HandlerFunc: ExampleHandler.Create},
	{Name: "ExampleReadAll", Methods: []string{"GET"}, Pattern: "/example", HandlerFunc: ExampleHandler.ReadAll},
	{Name: "ExampleReadOne", Methods: []string{"GET"}, Pattern: "/example/{id}", HandlerFunc: ExampleHandler.ReadOne},
	{Name: "ExampleUpdate", Methods: []string{"PUT"}, Pattern: "/example/{id}", HandlerFunc: ExampleHandler.Update},
	{Name: "ExampleDelete", Methods: []string{"DELETE"}, Pattern: "/example/{id}", HandlerFunc: ExampleHandler.Delete},
	{Name: "ExampleTask1", Methods: []string{"GET"}, Pattern: "/task/1/{name}", HandlerFunc: ExampleHandler.Task1},
	{Name: "ExampleTask2", Methods: []string{"GET"}, Pattern: "/task/2/{name}", HandlerFunc: ExampleHandler.Task2},
}

func init() {
	router.RouteApply(routes)
}
