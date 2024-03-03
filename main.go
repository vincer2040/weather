package main

import (
	"log"

	"github.com/vincer2040/weather/cmd/weather"
)

func main() {
    err := weather.Main()
    if err != nil {
        log.Fatalln(err)
    }
}
