package main

import ("fmt"
		"net/http")	//imports the liberary for http request handler

func index_handler(w http.ResponseWriter, r *http.Request){////handler function for the path "/" 
														////the first argument is for the writer the tpe is http Response writer which is used to write the response for any request that comes to "/" path
														////the second argument is for reading the http Request that comes at the path of "/"
					
						fmt.Fprintf(w, "Whoa, Go is neat!") // Fprintf formats based on whatever we specify here we specified w i.e: Response writer 
					}

func about_handler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "Expert server design by Asif")
}

func main(){
	http.HandleFunc("/", index_handler) //// http hanfler function
									   //// first parameter is the path for the handler
									   //// seond parameter is for the function that will handle the reques that comes to the path at first argument
	http.HandleFunc("/about/", about_handler )
	http.ListenAndServe(":8000",nil) //// http Listner and Server i.e : localhost:8000
									 //// first arguement is the port where server will start listening/taking http request
}									 ////second parameter is nil right now just let it we don't discuss it here