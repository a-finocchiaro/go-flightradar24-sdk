package main

import (
	"fmt"

	"github.com/a-finocchiaro/adsb_flightradar_top10/fr24"
)

func main() {
	var tracked fr24.Fr24MostTrackedRes = fr24.GetFR24MostTracked()
	fmt.Println(tracked.Data[0])
}
