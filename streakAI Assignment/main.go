package main

import (
	"net/http"
)

type GuestMux struct {
	http.ServeMux
}

func main() {
	guestMux := NewGuestMux()
	http.ListenAndServe(":3001", guestMux)
}

func NewGuestMux() *GuestMux {
	var guestMux = &GuestMux{}
	guestMux.HandleFunc("/find-pairs", findPairs)

	return guestMux
}
