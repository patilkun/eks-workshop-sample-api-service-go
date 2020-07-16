package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

type response struct {
	Message string   `json:"message"`
	EnvVars []string `json:"env"`
}

func hello() string {
	var helloInput string = "Demo CI/CD Workshop today with everyone!!"
	return helloInput
}

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hi there, I love %s!", hello())

	res := &response{Message: hello()}

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		res.EnvVars = append(res.EnvVars, pair[0]+"="+pair[1])
	}
	sort.Strings(res.EnvVars)

	// Beautify the JSON output
	out, _ := json.MarshalIndent(res, "", "  ")

	// Normally this would be application/json, but we don't want to prompt downloads
	w.Header().Set("Content-Type", "text/plain")

	io.WriteString(w, string(out))

	fmt.Println("Http Call")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}