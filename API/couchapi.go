package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getapifunc(w http.ResponseWriter, r *http.Request) {
	//resp, err := http.Get("http://127.0.0.1:5984/_all_dbs")
	resp, err := http.Get("http://127.0.0.1:5984/baseball2/8b8b6e539c5f806fe5e8aa50a00033f6")
	if err != nil {
		log.Print(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Fprintf(w, string(body))
}
type posting struct {
	string "Name" 
}
	myname posting = {"asif"}
func postapifunc(w http.ResponseWriter, r *http.Request) {
	//resp, err := http.Get("http://127.0.0.1:5984/_all_dbs")
	//resp, err := http.Post("http://127.0.0.1:5984/baseball2/",url.Values{"key":{"Value"},"name":{"asif"}})
	resp, err := http.Post("http://127.0.0.1:5984/baseball2/",posting,myname)
	if err != nil {
		log.Print(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Fprintf(w, string(body))
}

func main() {
	http.HandleFunc("/getapi", getapifunc)
	http.HandleFunc("/postapi", postapifunc)
	http.ListenAndServe(":8000", nil)
}
