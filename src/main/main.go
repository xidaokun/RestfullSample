package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"../github.com/gorilla/mux"
)

func main()  {
	//启动一个RESTful服务
	http.HandleFunc("/listener", func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(responseWriter, "test http, %q", html.EscapeString(request.URL.Path))
	})

	http.HandleFunc("/", TodoIndex)

	http.HandleFunc("/show" , TodoShow)

	log.Fatal(http.ListenAndServe(":10808", nil))
}

func TodoIndex(responseWriter http.ResponseWriter, request *http.Request)  {
	fmt.Fprintln(responseWriter, "Welcome!")
}

func TodoShow(responseWriter http.ResponseWriter, request *http.Request)  {
	vars := mux.Vars(request)
	todoId := vars["todoId"]
	fmt.Fprintln(responseWriter, "Todo show", todoId)
}