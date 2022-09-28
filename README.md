# gobybit

gobybit is a [Go](http://golang.org/) module for using the [ByBit's](https://www.bybit.com/) Rest & Websocket API


### Documentation

Full API, examples, and implementation notes are available in the Go
documentation.

[![Go Reference](https://pkg.go.dev/badge/github.com/ginarea/gobybit.svg)](https://pkg.go.dev/github.com/ginarea/gobybit)

### Installation

    go get github.com/ginarea/gobybit@latest

### Import

    import "github.com/ginarea/gobybit"

### Example
```
package main

import "github.com/ginarea/gobybit"

func main() {
    client := gobybit.NewClient()
    client.InversePerpetual().ServerTime()
    client.UsdtPerpetual().ServerTime()
    client.InverseFutures().ServerTime()
    client.Spot().ServerTime()
    client.Spotv3().ServerTime()
    client.AccountAsset()
}
```
