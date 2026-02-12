// Package router contains internal implementations of app related things like dependency injection etc
package router

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	return &http.ServeMux{}
}

// //func (r *Router) Get(url string,handler )
