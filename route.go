package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Post struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var posts []Post

func init() {
	fmt.Println("main init func")
	posts = []Post{
		{1, "nehal"}}
}

func listPosts(w http.ResponseWriter, r *http.Request) {
	posts_json, err := json.Marshal(posts)
	w.Header().Set("Content-type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(posts_json)
}

func addPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var post Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{message:error in unmarshal request body}`))
		return
	}
	post.Id = len(posts) + 1
	posts = append(posts, post)
	posts_json, err := json.Marshal(posts)

	w.Write(posts_json)

}

func updatePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	w.Header().Set("Content-type", "application/json")
	post_id, _ := strconv.Atoi(mux.Vars(r)["post_id"])

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"error in decoding"}`))
		return
	}
	for index, postVal := range posts {
		if postVal.Id == post_id {
			post.Id = post_id
			posts[index] = post
		}
	}
	res, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"error in decoding"}`))
		return
	}
	w.Write(res)
}
func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	post_id, _ := strconv.Atoi(mux.Vars(r)["post_id"])
	for index, postVal := range posts {
		if postVal.Id == post_id {
			posts = append(posts[:index], posts[index+1:]...)
		}
	}
	res, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"error in decoding"}`))
		return
	}
	w.Write(res)
}
