package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)



// Global variables
var counter int       // Stores the current count
var mutex sync.Mutex  // Used to ensure thread-safe access to the counter

func main () {
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// This is to set up my route handlers
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/increment", handleIncrement)
	http.HandleFunc("/decrement", handleDecrement)
	http.HandleFunc("/value", handleValue)
	http.HandleFunc("/reset", handleReset)

	//This prints ther server start message
	fmt.Println("Server is running on http://localhost:8080")

	// Starts server and logs any errors
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//Serves main html page
func handleHome(w http.ResponseWriter, r *http.Request) {
	resetCounter()
	http.ServeFile(w, r, "index.html")
}

// handle increment, increase counter by 1, return new value
func handleIncrement(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	mutex.Unlock()
	fmt.Fprintf(w, strconv.Itoa(counter)) // sends new value as response
}

// handle decrement, decrease counter by 1, return new value
func handleDecrement(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter--
	mutex.Unlock()
	fmt.Fprintf(w, strconv.Itoa(counter)) // sends new value as response
}

// returns current value of counter
func handleValue(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, strconv.Itoa(counter))
}

func handleReset(w http.ResponseWriter, r *http.Request) {
    resetCounter()
    fmt.Fprintf(w, strconv.Itoa(counter))
}

func resetCounter() {
    mutex.Lock()
    counter = 0
    mutex.Unlock()
}




