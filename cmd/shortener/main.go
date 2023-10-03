package main

import (
	"fmt"
	"net/http"
	"strings"
)

var database = make(map[string]string, 10)
var counter = 0

func shortname(s string) string {
	counter++
	return "url" + fmt.Sprint(counter)
}

func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		r.ParseForm()
		url := r.FormValue("url")
		shortUrl := shortname(url)
		database[shortUrl] = url

		w.Write([]byte(shortUrl))
		w.WriteHeader(http.StatusCreated)

	} else if r.Method == http.MethodGet {

		arr := strings.Split(r.URL.Path, "/")
		shortUrl := arr[len(arr)-1]

		v, ok := database[shortUrl]
		if ok {
			w.Header().Set("Location", v)
			w.WriteHeader(http.StatusTemporaryRedirect)
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

func main() {

	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
