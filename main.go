package main

import (
	"fmt"
	"log"

	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/client"
	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/common"
	"github.com/a-finocchiaro/go-flightradar24-sdk/webrequest"
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

	// res, err := client.GetAirportBrief(requester, "LHR")
	// fmt.Println(err)
	// fmt.Println(res)

	// zoneres, err := client.GetZones(requester)
	// fmt.Println(zoneres)
	// my_str := "plsugin[]=some_str"
	// my_str += "&plugin[]=some_str2"
	// fmt.Println(my_str)

	// routeres, err := client.GetAirportRoutes(requester, "tus", "SAN")
	// fmt.Println(routeres)

	// details, err := client.GetAirportDetails(requester, "TUS", []string{"details"})
	// fmt.Println(len(details.Schedule.Arrivals.Data))

	// zones, err := client.GetFlightsInZone(requester, "52.567967,13.282644,2000")
	radius := client.GetBoundsByPoint(32.918559, -97.058446, 500*1000)
	fmt.Println(radius)

}
