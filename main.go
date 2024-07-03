package main

import (
	"log"

	"github.com/Mat12143/WTracker/weatherapi"
)

func main()  {

    api := weatherapi.CreateWeatherAPI(
        []weatherapi.City{
            { Name: "Bolzano", Lat: 46.4949, Lon: 11.3403},
        },
    )

    log.Println(api.RetrieveWeathers())

}
