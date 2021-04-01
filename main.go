package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const CRYPTOURL = "wss://socket.polygon.io/crypto"

var dataStore = make(map[time.Time]CryptoAggregate)

func main() {
	channel := flag.String("channel", "", "")
	apiKey := flag.String("apiKey", "", "")

	flag.Parse()

	if *channel == "" || *apiKey == "" {
		logrus.Fatal("Please specify both -channel and -apiKey when running this program!")
		os.Exit(1)
	}

	c, _, err := websocket.DefaultDialer.Dial(CRYPTOURL, nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	_ = c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("{\"action\":\"auth\",\"params\":\"%s\"}", *apiKey)))
	_ = c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("{\"action\":\"subscribe\",\"params\":\"%s\"}", "XT." + *channel)))

	// Buffered channel to account for bursts or spikes in data:
	chanMessages := make(chan []CryptoTrade, 10000)

	// Read messages off the buffered queue:
	go func() {
		for trades := range chanMessages {
			dataStore = AddToDatastore(dataStore, trades)
		}
	}()

	// Print out aggregate bar in 30 second intervals
	tasks := cron.New()
	tasks.AddFunc("*/30 * * * *", func() {
		timestamp := time.Now().Truncate(30 * time.Second)
		PrintAggregate(dataStore[timestamp])
	})
	tasks.Start()

	// As little logic as possible in the reader loop:
	for {
		_, content, err := c.ReadMessage()

		if err != nil {
			logrus.Info("1005 ERROR: ", err)
			return
		}

		var trades []CryptoTrade
		err = json.Unmarshal(content, &trades)

		if err != nil {
			logrus.Info(err)
			return
		}

		chanMessages <- trades
	}
}