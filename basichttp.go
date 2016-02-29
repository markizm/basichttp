package main

import (
	"fmt"
	//	"log"
	"net/http"
	//	"sort"
	//	"strconv"
	//	"strings"
)

// basic web page with search field. Will eventually have this in a separate file once html and jquery is written
const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Search Tool</title>
<body><h3>VIP Viewer</h3>
<p>Searching for something?</p>`
	form = `<form action="/" method="POST">
<label for="numbers">Enter text to search:</label><br />
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Search">
</form>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
	outFormat  = `<table border="1">
  <tr><th colspan="3">Results</th></tr>
  <tr><td>key one</td><td>value one</td></tr>
  <tr><td>Status</td><td>Online, Available</td></tr>
  <tr><td>IP</td><td>192.x.x.x</td></tr>
  <tr><td>Port</td><td>80/443</td></tr>
  <tr><td>LTM</td><td>F5</td></tr>
  <tr><td>Pool</td><td>vip_443</td></tr>
  <tr><td>Pool Member</td><td>blah01.prod.dc...</td><td>green</td></tr>
  <tr><td>Pool Member</td><td>blah02.prod.dc...</td><td>red</td></tr>
  <tr><td>Pool Member</td><td>blah03.prod.dc...</td><td>green</td></tr>
  </table>`
)

// will eventually use these structs to parse and gather vip info
type vipinfo struct {
	vipname bool
	status  string
	ip      bool
	port    int
	ltm     bool
	pool    bool
}

// not sure if this will be needed yet, depends on how formatting of pool members will be
type poolMems struct {
	fqdn   bool
	status string
}

// using 9001 as default port
func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":9001", nil)
}

// have to edit this when i have solid funcs written for vipinfo values
func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm() // Must be called before writing response
	fmt.Fprint(writer, pageTop, form)
	fmt.Fprint(writer, pageBottom)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else if err == nil {
		fmt.Fprint(writer, outFormat)
	}
}

// func processRequest(request *http.Request) ([]float64, string, bool) {
// var vip []float64
// if slice, found := request.Form["vip"]
// }
