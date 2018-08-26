package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"traefik-auth-test/lib/utils"
)

const authCookieName = "auth"
const authCookieValue = "secret"

type Message struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.PrintRequest(r)
		c, err := r.Cookie(authCookieName)
		if err != nil {
			log.Println("No cookie")
			utils.Error(w, http.StatusForbidden, "Forbidden")
			return
		}
		cookieValue := c.Value
		if cookieValue == authCookieValue {
			msg := Message{
				Message: "OK",
			}
			utils.Ok(w, msg)
		} else {
			utils.Error(w, http.StatusForbidden, "Forbidden")
		}
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		cookie := &http.Cookie{
			Name:     authCookieName,
			Value:    authCookieValue,
			HttpOnly: false,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	})
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		cookie := &http.Cookie{
			Name:     authCookieName,
			Value:    "",
			HttpOnly: false,
			Path:     "/",
			Expires:  time.Unix(0, 0),
		}
		http.SetCookie(w, cookie)
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
