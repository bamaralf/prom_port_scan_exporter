package main

import (
	"encoding/json"
	"fmt"
    "net/http"
)


func healthCheck(w http.ResponseWriter, req *http.Request) { 
	output,_ := json.Marshal( "200: {Status: OK}" )
    fmt.Println(w, output)
}