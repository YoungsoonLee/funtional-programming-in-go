package interfaces

import "net/http"

type Api struct {
	Handler func(res http.ResponseWriter, req *http.Request)
	Url     string
}
