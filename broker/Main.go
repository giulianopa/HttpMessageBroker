/*
 * Copyright (C) 2018 Giuliano Pasqualotto (github.com/giulianopa)
 * This code is licensed under MIT license (see LICENSE.txt for details)
 */
package main

import (
		"log"
		"net/http"
)

func main() {
		router := NewRouter()

		log.Fatal(http.ListenAndServe(":8080", router))
}
