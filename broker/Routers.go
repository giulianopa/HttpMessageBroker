/*
 * Copyright (C) 2018 Giuliano Pasqualotto (github.com/giulianopa)
 * This code is licensed under MIT license (see LICENSE.txt for details)
 */
package main

import (
		"net/http"
		"github.com/gorilla/mux"
)

type Route struct {
		Name        string
		Method      string
		Pattern     string
		HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

		router := mux.NewRouter().StrictSlash(true)
		for _, route := range routes {
				router.
						Methods(route.Method).
						Path(route.Pattern).
						Name(route.Name).
						Handler(route.HandlerFunc)
		}

		return router
}

// Endpoints
var routes = Routes{
		Route{
				"Message",
				"POST",
				"/message",
				ReceiveMessage,
		},
}
