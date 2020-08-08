package front

import (
	"net/http"
	"github.com/david-ch/otus-architect/front-service/front/controller"
	"github.com/gorilla/mux"
)

// RegisterFrontRoutes registers all the routes of the package
func RegisterFrontRoutes(r *mux.Router) {
	regPath(r, "/formattedRandomNumber", "/", "GET", controller.OnGetFormattedRandomNumber)
}

func regPath(r *mux.Router, path string, metricsName string, httpMethod string, handler func(http.ResponseWriter, *http.Request)) *mux.Route {
	return r.HandleFunc(path, handler).Methods(httpMethod)
}
