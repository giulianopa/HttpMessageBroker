package main

type Message struct {
    Name      string    `json:"name"`
    Image     string    `json:"image"`
}

type Messages []Message
