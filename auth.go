package main

import (
	"net/http"
)

func checkAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		user, pass , ok := r.BasicAuth()
		if ok && user == "test" && pass == "pass123" {
			handler( w, r )
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			w.Header().Set("WWW-Authenticate", "Basic realm=\"MY REALM\"")
		}
	}
}

func main() {
	http.HandleFunc("/", checkAuth( handleIndex ))

	http.ListenAndServe(":2424", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request){

}

