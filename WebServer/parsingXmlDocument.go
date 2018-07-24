package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil" //library to read the bytes data from body of the response
	"net/http"  //library to trigger http get request
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

// type SitemapIndex struct {
// 	Locations []Location `xml:"sitemap"`
// }

// type Location struct{
// 	Loc string `xml:"loc"`
// }

// func (l Location) String()  string {
// 	return fmt.Sprintf(l.Loc)
// }

// [5] type == array //e.g: [5] int
// [] type == slice //e.g: [] int

func main() {
	// resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml") //get request _ is used for the variable that we donot intend to use otherwise go will throw an error
	// bytes, _ := ioutil.ReadAll(resp.Body) //reading byte from body of response byte data
	// stringBody := string(bytes) //typecasting byte data into string
	// fmt.Println(stringBody) //print string body
	// resp.Body.Close() 		// Close body of byte data

	// var s SitemapIndex //s variable declared
	// byte,_ := ioutil.ReadAll(resp.Body)
	// xml.Unmarshal(bytes,&s)
	// // xml.Unmarshal(bytes, &s) //
	// //fmt.Println(s.Locations)
	// for _,Location := range s.Locations{
	// 	fmt.Printf("\n%s", Location)
	// }

	var s SitemapIndex
	var n News

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")

	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	news_map := make(map[string]NewsMap)
	for _, Location := range s.Locations {

		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		for idx, _ := range n.Titles {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}
	for idx, data := range news_map {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)

	}

}
