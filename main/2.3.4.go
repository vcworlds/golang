package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func myGet(w http.ResponseWriter, r *http.Request) {
	arg1 := r.URL.Query().Get("arg1")
	arg2 := r.URL.Query().Get("arg2")
	response := fmt.Sprintf("接收到参数，分别为: arg1: %s, arg2: %s", arg1, arg2)
	w.Write([]byte(response))
}

func myPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var getData map[string]interface{}
	err := decoder.Decode(&getData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := fmt.Sprintf("接收到请求数据，内容为:%v", getData)
	w.Write([]byte(response))
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/my_get", myGet)
	http.HandleFunc("/my_post", myPost)
	http.ListenAndServe(":9090", nil)
}
