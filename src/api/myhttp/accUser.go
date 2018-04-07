package myhttp

import (
	"api/db"
	"encoding/json"
	"net/http"
)

func accGetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	u, err := db.SelectUser(id)
	if err == db.ErrNotFound {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	showUser(w, u)
}

func accSetUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u db.User
	if err := decoder.Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	if err := u.Insert(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	showUser(w, &u)
}

func accUpdateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u db.User
	if err := decoder.Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	if err := u.Update(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	showUser(w, &u)
}

func accDeleteUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u db.User
	if err := decoder.Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	if err := u.Delete(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	showUser(w, &u)
}

func showUser(w http.ResponseWriter, u *db.User) {
	w.Header().Set("Content-Type", "application/json")

	b, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}
