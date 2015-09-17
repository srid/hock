package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"strconv"
)

func oops(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERR: %v\n", err)
		os.Exit(1)
	}
}

func drainGateway(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	if name == "" {
		oops(fmt.Errorf("no name"))
	}
	drn := getOrCreateDrain(name)
	drn.LogsHandler(w, r)
}

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		oops(err)
	}
	addr := fmt.Sprintf("0.0.0.0:%d", port)

	router := httprouter.New()
	router.POST("/logs/:name", drainGateway)

	err = http.ListenAndServe(addr, router)
	oops(err)
}
