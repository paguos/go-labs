package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", process_request)
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func process_request(w http.ResponseWriter, r *http.Request) {
	names := r.URL.Query()["name"]
	m := map[string][]string{"names": names}
	enc := json.NewEncoder(w)
	enc.Encode(m)
}
