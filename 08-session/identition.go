package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type uuidSession struct {
}

func (uuidSession) executeMain() {
	http.Handle("/", http.HandlerFunc(home))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8089", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil { // no cookie available
		id := uuid.New()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true, // attach to request header when connection is https
			HttpOnly: true, // can not be accessed with javscript, only server can accessed
		}
		http.SetCookie(w, cookie)
	}
	fmt.Fprintf(w, "Session: %+v\n", cookie)
}
