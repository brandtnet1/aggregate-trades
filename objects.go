package main

import (
	"time"
)

type CryptoTrade struct {
	Ev   string  `json:"ev"`   // The event type.
	Pair string  `json:"pair"` // The crypto pair.
	P    float64 `json:"p"`    // The price.
	T    int64   `json:"t"`    // The Timestamp in Unix MS.
	S    float64 `json:"s"`    // The size.
	C    []int   `json:"c"`    // The conditions. 0 (or empty array): empty 1: sellside 2: buyside
	I    string  `json:"i"`    // The ID of the trade (optional).
	X    int     `json:"x"`    // The crypto exchange ID. See Crypto Exchanges for a list of exchanges and their IDs.
	R    int64   `json:"r"`    // The timestamp that the tick was recieved by Polygon.
}

type CryptoAggregate struct {
	Timestamp time.Time
	Open      float64
	Close     float64
	High      float64
	Low       float64
	Volume    int
	BeginTime time.Time
	EndTime   time.Time
}