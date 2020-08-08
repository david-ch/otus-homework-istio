package back

import (
	"net/http"
	"github.com/david-ch/otus-architect/back-service/back/controller"
	"github.com/gorilla/mux"
)

// RegisterBackRoutes registers all the routes of the package
func RegisterBackRoutes(r *mux.Router) {
	regPath(r, "/randomNumber", "/", "GET", controller.OnGetRandomNumber)
}

func regPath(r *mux.Router, path string, metricsName string, httpMethod string, handler func(http.ResponseWriter, *http.Request)) *mux.Route {
	return r.HandleFunc(path, handler).Methods(httpMethod)
}
