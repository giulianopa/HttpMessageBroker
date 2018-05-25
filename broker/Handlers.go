package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func ReceiveMessage(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)

  	var m Message
  	err := decoder.Decode(&m)
    if err != nil {
  		panic(err)
  	}

  	fmt.Println(m.Name)
}
