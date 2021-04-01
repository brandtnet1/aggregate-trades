package main

import (
	"fmt"
	"time"
)

func GetTimestamp(timestamp time.Time) time.Time {
	return timestamp.Truncate(30 * time.Second).Add(30 * time.Second)
}

func AddToDatastore(dataStore map[time.Time]CryptoAggregate, cryptoTrades []CryptoTrade) map[time.Time]CryptoAggregate {
	timestamp := time.Time{}

	for _, cryptoTrade := range cryptoTrades {
		tradeTime := time.Unix(cryptoTrade.T / 1000, 0)

		if tradeTime.After(time.Now().Add(-time.Hour * 1)) {
			timestamp = GetTimestamp(tradeTime)

			aggregate := dataStore[timestamp]
			aggregate.Timestamp = timestamp
			aggregate.Volume += 1

			if tradeTime.Before(dataStore[timestamp].BeginTime) || dataStore[timestamp].BeginTime.Equal(time.Time{}) {
				aggregate.Open = cryptoTrade.P
				aggregate.BeginTime = tradeTime
			}
			if tradeTime.After(dataStore[timestamp].EndTime){
				aggregate.Close = cryptoTrade.P
				aggregate.EndTime = tradeTime
			}

			if aggregate.High < cryptoTrade.P {
				aggregate.High = cryptoTrade.P
			}
			if aggregate.Low > cryptoTrade.P || aggregate.Low == 0 {
				aggregate.Low = cryptoTrade.P
			}

			dataStore[timestamp] = aggregate

			// The case where we receive a trade with a timestamp before now, print that updated aggregate
			if !timestamp.Equal(time.Time{}) && GetTimestamp(time.Now()).After(timestamp) {
				PrintAggregate(dataStore[timestamp])
			}
		}
	}

	return dataStore
}

func PrintAggregate (aggregate CryptoAggregate) {
	if !aggregate.Timestamp.Equal(time.Time{}) {
		fmt.Printf("%v - open: $%.2f, close: $%.2f, high: $%.2f, low: $%.2f, volume: %v\n", aggregate.Timestamp.Format("15:04:05"), aggregate.Open, aggregate.Close, aggregate.High, aggregate.Low, aggregate.Volume)
	}
}
