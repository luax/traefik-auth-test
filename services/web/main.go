package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"traefik-auth-test/lib/utils"
)

const indexPage = `
  <h1>Index</h1>
  <hr>
  <a href="/login">Login</a>
  <a href="/search">Search</a>
  <script src="assets/main.js"></script>
`

const loginPage = `
  <h1>Login</h1>
  <a href="/">Start</a>
  <hr>
  <div id="text"></div>
  <button id="login">Login</button>
  <button id="logout">Logout</button>
  <script src="assets/main.js"></script>
`

const searchPage = `
  <h1>Search</h1>
  <a href="/">Start</a>
  <hr>
  <div id="text"></div>
  <hr>
  <input id="search-input" type="text" value="Hello">
  <button id="search">Search</button>
  <script src="assets/main.js"></script>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.PrintRequest(r)
		fmt.Fprintf(w, indexPage)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		utils.PrintRequest(r)
		fmt.Fprintf(w, loginPage)
	})
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		utils.PrintRequest(r)
		fmt.Fprintf(w, searchPage)
	})
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
