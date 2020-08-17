package main

import (
	"fmt"
    "net/http"
)


func healthCheck(w http.ResponseWriter, req *http.Request) { 
    fmt.Fprintf(w, "200: {Status: OK}")
}