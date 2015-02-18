package main

type Message struct {
	Text string `json:"message"`

	To string `json:"msisdn"`

	Response int
}
