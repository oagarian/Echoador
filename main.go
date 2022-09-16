package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func handler(w http.ResponseWriter, req *http.Request) {
	
	values := req.URL.Query().Get("echo")
	split_values := strings.Split(values, "-")
	value_one := split_values[0]
	value_two := split_values[1]
	
	repeat, err := strconv.Atoi(value_one)
	if err != nil {
		log.Fatal(err)
	}
	var str bytes.Buffer
	string_one := "\n"
	for i := 0; i < repeat; i++{
		str.WriteString(value_two+string_one)
	}
	fmt.Fprintf(w, "%s", str.String())



}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
