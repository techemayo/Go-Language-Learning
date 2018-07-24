package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getapifunc(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://127.0.0.1:5984/_all_dbs")
	if err != nil {
		fmt.Fprintf(w, err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Fprintf(w, body)
}

func main() {
	http.HandleFunc("/getapi", getapifunc)
}
