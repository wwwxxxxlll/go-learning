package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"

	jq "github.com/savaki/jq"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		fmt.Fprintf(w, "\nURL.Path = %s\nhere is a response body, with id %s", r.URL.Path, id)
	})
	http.HandleFunc("/test1", func(w http.ResponseWriter, r *http.Request) {
		body, rerr := ioutil.ReadAll(r.Body)
		if rerr != nil {
			panic(rerr)
		}
		defer r.Body.Close()
		op, operr := jq.Parse(".title")
		if operr != nil {
			panic(operr)
		}
		id, aerr := op.Apply(body)
		if aerr != nil {
			fmt.Println(string(body))
			panic(aerr)
		}
		w.Write(([]byte)("{\"title\":"))
		w.Write(id)
		w.Write(([]byte)("}"))
	})
	http.ListenAndServe("localhost:8502", nil)
}
