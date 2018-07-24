package main

import (
	"time"

	"github.com/rhinoman/couchdb-go"

)

// type TestDocument struct {
// 	Title string
// 	Note  string
// }

func main() {

	// var timeout = time.Duration(500 * time.Millisecond)
	// conn, err := couchdb.NewConnection("127.0.0.1", 5984, timeout)
	// auth := couchdb.BasicAuth{Username: "root", Password: "asif4106"}
	// err = conn.CreateDB("baseball2", &auth)
	// if err != nil {
	// 	println("Database Created")
	// }
	//var doc TestDocument
	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("127.0.0.1", 5984, timeout)
	auth := couchdb.BasicAuth{Username: "root", Password: "asif4106"}
	db := conn.SelectDB(" 
	
	", &auth)
	doc2, err := db.Read("346677", )
	if err != nil {
		println(err)
	} else {
		println(doc)
	}
	// theDoc := TestDocument{
	// 	Title: "My Document",
	// 	Note:  "This is a note",
	// }

	// //theId := genUuid() //use whatever method you like to generate a uuid
	// theId := "346677"

	// //The third argument here would be a revision, if you were updating an existing document
	// rev, err := db.Save(theDoc, theId, "")
	// //If all is well, rev should contain the revision of the newly created
	// //or updated Document
	// if err != nil {
	// 	println("doc inserted", rev)
	// } else {

	// 	println(err)

	// }
}
