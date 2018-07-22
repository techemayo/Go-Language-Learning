package main

import ("fmt"
		"net/http"
		"html/template")

type NewsAggPage struct {
	Title string
	News string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request){
	p:=NewsAggPage{Title: "The Amazing News", News: "Some News"}
	t,_ := template.ParseFiles("basichtmltemplate.html")
	fmt.Println(t.Execute(w,p))
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Whoa, Go is Neat!</h1>")

}

func main() {
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000",nil)
}