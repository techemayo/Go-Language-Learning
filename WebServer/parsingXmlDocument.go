package main

import ("fmt"
		"net/http" //library to trigger http get request
		"io/ioutil" //library to read the bytes data from body of the response 
		"encoding/xml")

type SitemapIndex struct {
	Locations []Location `xml:"sitemap>loc"`
}

type News struct{
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>Keywords"`
	Location []string `xml:"url>loc"`
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


func main(){
	// resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml") //get request _ is used for the variable that we donot intend to use otherwise go will throw an error
	// bytes, _ := ioutil.ReadAll(resp.Body) //reading byte from body of response byte data
	// stringBody := string(bytes) //typecasting byte data into string
	// fmt.Println(stringBody) //print string body
	// resp.Body.Close() 		// Close body of byte data

	// var s SitemapIndex //s variable declared
	byte,_ := iotutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes,&s)
	// xml.Unmarshal(bytes, &s) //
	// //fmt.Println(s.Locations)
	// for _,Location := range s.Locations{
	// 	fmt.Printf("\n%s", Location)
	// }

	var s SitemapIndex
	var n News
	resp,_ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	for _,Location :=range s.Location{
		resp,_:=ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes,&n)
	}
}
