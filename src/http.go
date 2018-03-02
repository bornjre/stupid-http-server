package core

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Run() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	fileserver := http.FileServer(http.Dir(dir))

	http.HandleFunc("/", func(h http.ResponseWriter, r *http.Request) {
		log.Println("REQUEST: ", r.URL, "->", r.Method)

		fileserver.ServeHTTP(h, r)
	})

	http.ListenAndServe(":8080", nil)
}
