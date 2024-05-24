package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"

	"github.com/deatil/go-encoding/base62"
	"github.com/ssdb/gossdb/ssdb"
)

// Преобразует URL в 8-символьный код в формате [A-Za-z0-9]{8}
func hashUrl(url string) string {
	sum := sha256.Sum256([]byte(url))
	b62 := base62.StdEncoding.EncodeToString(sum[:])
	return b62[0:8]
}

func main() {

	db, err := ssdb.Connect("127.0.0.1", 8888)
	if err != nil {
		fmt.Println("ssdb.Connect() error")
		os.Exit(1)
	}

	http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {

		url := r.URL.Query().Get("url")
		hash := hashUrl(url)

		fmt.Fprintln(w, hash)
		db.Set(hash, url)
	})

	http.HandleFunc("/s/{code}", func(w http.ResponseWriter, r *http.Request) {

		var url interface{}
		url, err = db.Get(r.PathValue("code"))
		http.Redirect(w, r, fmt.Sprintf("%s", url), http.StatusFound)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Index page\n")
		fmt.Fprintf(w, "Try for save url: \n http://127.0.0.1:8080/a?url=http://google.com/?q=golang \n")
		fmt.Fprintf(w, "Try for get url:  \n http://127.0.0.1:8080/s/cDQIQzgt \n")
	})
	http.ListenAndServe(":8080", nil)
}
