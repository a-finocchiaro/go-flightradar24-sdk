package main

import (
	"fmt"
	"log"

	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/client"
	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/models/common"
	"github.com/a-finocchiaro/adsb_flightradar_top10/webrequest"
)

func main() {
	var requester common.Requester = webrequest.SendRequest
	tracked, err := client.GetFR24MostTracked(requester)

	if err != nil {
		log.Fatalln("Something bad happened")
	}

	fmt.Println(tracked.Data[0])

	// give me a random flight link
	// fr24.GetRandomFlight(requester)

	// var myFeed fr24.Fr24FeedData
	// fr24.GetFlights(requester, &myFeed)

	// fmt.Println(myFeed)

	// res, err := fr24.GetAirlineLogoCdn(requester, "WN", "SWA")
	// fmt.Println(res)

	res, err := client.GetAirportBrief(requester, "LHR")
	fmt.Println(err)
	fmt.Println(res)
	// my_str := "plsugin[]=some_str"
	// my_str += "&plugin[]=some_str2"
	// fmt.Println(my_str)
}
