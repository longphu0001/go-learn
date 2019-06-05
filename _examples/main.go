package main

import (
	"fmt"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprintf(w, "Hi there, %s!", r.URL.Path[1:])
// 	fmt.Fprintf(w, "<!--\n"+html.EscapeString(spew.Sdump(w))+"\n-->")
// }

func main() {
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)

	fmt.Println("======================================")

	dataTest := dataTest()

	inDumpWithSpew(dataTest)
	inDumpWithPP(dataTest)

	fmt.Println("======================================")
}
