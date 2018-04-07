package myhttp

import "net/http"

type handleerFn func(http.ResponseWriter, *http.Request)

func switchMetod(handlerGet, handlerPost, handlerPut, handlerDel handleerFn) handleerFn {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlerGet(w, r)
		case "POST":
			handlerPost(w, r)
		case "PUT":
			handlerPut(w, r)
		case "DELETE":
			handlerDel(w, r)
		default:
			http.Error(w, "Method not correct", http.StatusBadRequest)
		}
	}
}
