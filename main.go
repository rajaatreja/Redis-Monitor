/*
Author- Raja Atreja
*/

package main

import (
	"fmt"
	"log"
	"net/http"

	"./websocket" //private package
	"github.com/gorilla/mux"
)

//defining localhost homepage with simple http router connection
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1><center>Redis Monitor</center></h1>")
}

func main() {

	//Initialising new http router
	r := mux.NewRouter()
	//routing two pages 1.Homepage & 2.WebSocket page
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/ws", websocket.WsEndpoint).Methods("GET")

	/*
		To access WebSocket and print messages from server type
		this in console of browser =>
		var ws = new WebSocket("ws://localhost:8080/ws");
		ws.addEventListener("message", function(e) {console.log(e.data);})

	*/

	// open => localhost:8080
	log.Fatal(http.ListenAndServe(":8080", r))
	//To stop server => ctrl+c
}
