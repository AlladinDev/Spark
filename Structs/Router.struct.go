package structs

import "net/http"

type Router struct {
	Mux *http.ServeMux
}
