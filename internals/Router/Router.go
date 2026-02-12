// Package Router contains internal implementations of app related things like dependency injection etc
package Router

import (
	"net/http"

	Structs "github.com/AlladinDev/Spark/Structs"
)

func NewRouter() *Structs.Router {
	return &Structs.Router{
		Mux: http.NewServeMux(),
	}
}

// //func (r *Router) Get(url string,handler )
