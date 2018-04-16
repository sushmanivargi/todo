package route

import (
	"net/http"

	"app/controller"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)

	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"Show",
		"GET",
		"/tasks/{taskId}",
		controller.Show,
	},
	Route{
		"Create",
		"POST",
		"/tasks",
		controller.Create,
	},
	// 	Route{
	// 		"Update",
	// 		"PUT",
	// 		"/tasks/{taskId}",
	// 		controller.Create,
	// 	},
	Route{
		"Destroy",
		"DELETE",
		"/tasks/{taskId}",
		controller.Destroy,
	},
}
