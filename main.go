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
	r.HandleFunc("/posts", listPosts).Methods("Get")
	r.HandleFunc("/posts", addPost).Methods("Post")
	r.HandleFunc("/posts/{post_id}", updatePost).Methods("Put")
	r.HandleFunc("/posts/{post_id}", deletePost).Methods("Delete")
	http.ListenAndServe(":9000", r)

}
