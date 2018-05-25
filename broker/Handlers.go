package main

import (
    "encoding/json"
    "encoding/base64"
    "bytes"
    "fmt"
    "os"
    "bufio"
    "strings"
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

    // Decode Base64 image
    stream := base64.NewDecoder(base64.StdEncoding, strings.NewReader(m.Image))
  	//img, _, err := image.Decode(reader)
  	//if err != nil {
  	//	panic(err)
  	//}

    // TEST:Write to file.
    buf := new(bytes.Buffer)
    buf.ReadFrom(stream)
    fo, err := os.Create("img.png")
    if err != nil {
      panic(err)
    }
    fw := bufio.NewWriter(fo)
    fw.Write(buf.Bytes())
}
