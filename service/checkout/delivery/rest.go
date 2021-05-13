package delivery

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lonysutrisno/tokenomy/pkg"
	"github.com/lonysutrisno/tokenomy/service/checkout/usecase"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"/", "GET", GetData},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
	return router
}

func GetData(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query().Get("id")

	res, code, err := usecase.GetData(params)
	if err != nil {
		pkg.JSONResponse(w, code, nil, err)
		return
	}
	pkg.JSONResponse(w, 200, res, nil)
}
