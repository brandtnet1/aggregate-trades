#Polygon.io Websockets

##Aggregation with out of order trades

When it comes to the stock market, aggregations are key. Polygon provides 1-minute aggregates through our websocket
server, but for this exercise, imagine you’re writing a trading algorithm that only works with 30-second aggregates.

Write a program that connects to Polygon’s websocket server, subscribes to the trades feed,
and outputs 30-second aggregates for a given ticker.

For the purposes of this exercise, an aggregate contains the following information:
Ticker symbol
Open price: the price of the first trade for the ticker in the aggregate window
Close price: the price of the last trade for the ticker in the aggregate window
High price: the highest price traded for the ticker during the aggregate window
Low price: the lowest price traded for the ticker during the aggregate window
Volume: the total volume of all the trades for the ticker in the aggregate window
Time: the timestamp of the start of the aggregate window
We can assume the end time of the aggregate window will always be 30 seconds after the start time, for this exercise.

We’ll provide you with an API key with permissions to connect to Polygon’s delayed websocket server (delayed.polygon.io/stocks).

Note: the key will also have permissions to connect to Polygon’s crypto currencies websocket server (socket.polygon.io/crypto).
Since crypto currencies trade 24/7, you may find it easier to test with the crypto server. Feel free to support only crypto
tickers, only stocks tickers, or both (but please know that we don’t care if you choose to support just one type of ticker!).

You can find documentation for Polygon’s websocket server here (https://polygon.io/docs/websockets/getting-started)

Here’s an example of what running this program might look like:

```shell
$ ./aggregate-trades -channel BTC-USD -apiKey <API-KEY> 
20:31:30 - open: $59049.24, close: $59039.40, high: $59051.75, low: $59009.35, volume: 23
20:32:00 - open: $59039.40, close: $59008.28, high: $59050.00, low: $58988.41, volume: 188
20:32:30 - open: $59043.84, close: $59058.59, high: $59070.96, low: $58992.66, volume: 167
20:33:00 - open: $59058.58, close: $59037.75, high: $59085.83, low: $58992.66, volume: 369
20:33:30 - open: $59073.50, close: $59043.84, high: $59080.29, low: $58992.00, volume: 188
```
*OR*
```shell
go build
/aggregate-trades.exe -channel BTC-USD -apiKey <API-KEY>
```
*OR*
```shell
go run . -channel BTC-USD -apiKey <API-KEY>
```

Important Note:
It’s possible for trades to be reported out of order, and sometimes these trades can be reported more than 30 seconds
after their execution time. If this happens, your program should be able to update an aggregate bar from the past and
output an updated bar. For example, if a trade executed at 13:00:01 comes through the websocket at 13:01:40, that trade
should update the aggregate bar that spans 13:00:00 - 13:00:30, and your program should output that bar again.

For our purposes, let’s say we only care about out of order trades if they’re reported within 1 hour of when they were
executed. Meaning if a trade executed at 13:00:01 but comes through the websocket at 14:12:00, we can disregard it and
not print an updated aggregate bar.