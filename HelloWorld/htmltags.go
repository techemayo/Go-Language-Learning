package main

import ("fmt"
		"net/http")

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Whoa, Go is Neat!</h1>")

}

func func main() {
	http.HandleFunc("/",indexHandler)
	http.ListenAndServe(":8000",nil)
}