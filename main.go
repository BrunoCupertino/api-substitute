package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
)

type valueFlags []string

func (v *valueFlags) String() string {
	return ""
}

func (v *valueFlags) Set(value string) error {
	*v = append(*v, value)
	return nil
}

var (
	port        string
	routes      valueFlags
	responses   valueFlags
	responseMap map[string]string
)

func main() {
	flag.StringVar(&port, "port", "1010", "port to listen to")
	flag.Var(&routes, "routes", "list of routes")
	flag.Var(&responses, "responses", "list of responses for routes")

	flag.Parse()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	fmt.Println("route /ping registered")

	responseMap = make(map[string]string, len(routes))

	for i, ep := range routes {
		http.HandleFunc(ep, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			rb, _ := httputil.DumpRequest(r, true)
			reqStr := string(rb)

			fmt.Printf("*** NEW REQUEST\n %s\n***\n", reqStr)

			fmt.Fprintf(w, responseMap[r.RequestURI])
		})

		responseMap[ep] = responses[i]

		fmt.Printf("route %s registered\n", ep)
	}

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
