package weatherapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type WeatherAPI struct {
    Cities []City
}

type City struct {
    Name string
    Lat float32
    Lon float32
}

type APIResponse struct {
    Weather []struct {
        Id int `json:"id"`
    } `json:"weather"`
    Main struct {
        Temp float32 `json:"temp"`
        Feels_like float32 `json:"feels_like"`
    } `json:"main"`
    Name string `json:"name"`
}

type Result struct {
    Name string
    WeatherID int
    Temperature float32
    FeelsLike float32
}

func CreateWeatherAPI(c []City) (*WeatherAPI) {
    return &WeatherAPI{Cities: c}
}

func (w *WeatherAPI) RetrieveWeathers() ([]Result) {

    godotenv.Load()

    var results []Result

    for _, c := range w.Cities {

        url := "https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&lang=it&units=metric&appid=%s"
        url = fmt.Sprintf(url, c.Lat, c.Lon, os.Getenv("OPEN_METEO_API_KEY"))

        log.Println(url)

        resp, err := http.Get(url)

        if err != nil {
            log.Println(err)
            continue
        }

        defer resp.Body.Close()

        body, err := io.ReadAll(resp.Body)

        if err != nil {
            log.Println(err)
            continue
        }

        if resp.StatusCode != 200 {
            log.Println(resp.StatusCode)
            log.Println(string(body))
            continue
        }

        var response APIResponse

        err = json.Unmarshal(body, &response)
        
        if err != nil {
            log.Println(err)
            continue
        }

        if c.Name != response.Name {
            log.Println("City provided and API returned city are not the same. Maybe Wrong Coordinates??")
            continue
        }

        results = append(results, Result{ Name: c.Name, WeatherID: response.Weather[0].Id, Temperature: response.Main.Temp, FeelsLike: response.Main.Feels_like })

    }
    return results
}
