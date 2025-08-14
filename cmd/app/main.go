package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vugsk/CurrencyExchangerProjectGoLang/internal/app"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, World! Method: %s, URL: %s", r.Method, r.URL.Path)
		if err != nil {
			return
		}
	})

	app.Fu()

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
