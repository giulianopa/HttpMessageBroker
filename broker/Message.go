/*
 * Copyright (C) 2018 Giuliano Pasqualotto (github.com/giulianopa)
 * This code is licensed under MIT license (see LICENSE.txt for details)
 */
package main

// Basic message packet
type Message struct {
		Name      string    `json:"name"`
		Image     string    `json:"image"`
}

type Messages []Message
