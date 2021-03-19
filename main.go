package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`welcome nehal`))
	})
	ps := r.PathPrefix("/posts").Subrouter().StrictSlash(true)
	ps.HandleFunc("/", listPosts).Methods("Get")
	ps.HandleFunc("/", addPost).Methods("Post")
	ps.HandleFunc("/{post_id}", updatePost).Methods("Put")
	ps.HandleFunc("/{post_id}", deletePost).Methods("Delete")
	http.ListenAndServe(":9000", r)

}
