package example

import "github.com/haifahrul/go-playground/internal/router"

var routes = []router.Route{
	{Name: "ExampleGet", Methods: []string{"GET"}, Pattern: "/example", HandlerFunc: ExampleHandler.Get},
	// {Name: "ExampleGetByID", Methods: []string{"GET"}, Pattern: "/example/{id}", HandlerFunc: ExampleHandler.GetByID},
	// {Name: "ExamplePost", Methods: []string{"POST"}, Pattern: "/example", HandlerFunc: ExampleHandler.Post},
}

func init() {
	router.RouteApply(routes)
}
