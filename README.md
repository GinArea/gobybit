# gobybit

gobybit is a [Go](https://go.dev) module for using the [ByBit's](https://bybit.com) Rest & Websocket API


### Documentation

Full API, examples, and implementation notes are available in the Go
documentation.

[![Go Reference](https://pkg.go.dev/badge/github.com/ginarea/gobybit.svg)](https://pkg.go.dev/github.com/ginarea/gobybit)

### Installation

    go get github.com/ginarea/gobybit@latest

### Import

    import "github.com/ginarea/gobybit/bybitv5"

### Example
```
package main

import (
    "fmt"
    "reflect"

    "github.com/ginarea/gobybit/bybitv5"
)

func main() {
    apiKey := "XXXXX"
    apiSecret := "XXXXX"

    client := bybitv5.NewClient().WithAuth(apiKey, apiSecret)

    responce(client.GetInstruments(bybitv5.GetInstruments{
        Category: bybitv5.Inverse,
        Symbol:   "BTCUSD",
    }))

    responce(client.GetTickers(bybitv5.GetTickers{
        Category: bybitv5.Inverse,
        Symbol:   "ETHUSD",
    }))

    responce(client.GetOrderbook(bybitv5.GetOrderbook{
        Category: bybitv5.Spot,
        Symbol:   "ETHUSDT",
    }))
}

func responce[T any](r bybitv5.Response[T]) {
    fmt.Printf("%v: %+v\n", reflect.TypeOf(r), r)
}
```
