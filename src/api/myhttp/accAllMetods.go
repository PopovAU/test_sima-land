package myhttp

import (
	"net/http"
)

func accAllMetods(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		GET    /user
		POST   /user
		PUT    /user
		DELETE /user
		`))
}
