package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"log"
	"errors"
)

func main() {
	mw := multiWeatherProvider{
		weatherUnderground{apiKey: "noapikey"},
		openWeatherMap{},
	}

	http.HandleFunc(
		"/weather/",
		func(w http.ResponseWriter, r *http.Request) {
			begin := time.Now()
			city := strings.SplitN(r.URL.Path, "/", 3)[2]

			temp, err := mw.temperature(city)
			if err != nil {
				log.Println("no result")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")

			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"city": city,
					"temp": temp,
					"took": time.Since(begin).String(),
				},
			)
		},
	)

	http.ListenAndServe(":8080", nil)
}

type multiWeatherProvider []weatherProvider

func (w multiWeatherProvider) temperature(city string) (float64, error) {
	var err error
	sum := 0.0
	nprov := 0
	for _, provider := range w {
		var k float64
		k, err = provider.temperature(city)
		if err != nil {
			//return 0, err
		} else {
			sum += k
			nprov++
		}
	}
	if nprov == 0 {
		return 0, err
	}
	return sum / float64(nprov), nil
	//return sum / float64(len(w)), nil
}

//func temperature(city string, providers ...weatherProvider) (float64, error) {
//    sum := 0.0
// 
//    for _, provider := range providers {
//        k, err := provider.temperature(city)
//        if err != nil {
//            return 0, err
//        }
// 
//        sum += k
//    }
// 
//    return sum / float64(len(providers)), nil
//}

type weatherProvider interface {
	temperature(city string) (float64, error) // in Kelvin, naturally
}


type openWeatherMap struct{}

func (w openWeatherMap) temperature(city string) (float64, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=????????????????????????????????&q=" + city)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	var d struct {
		Main struct {
			Kelvin float64 `json:"temp"`
		} `json:"main"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return 0, err
	}
	if d.Main.Kelvin == 0.0 {
		log.Printf("openWeatherMap: %s: not available", city)
		return d.Main.Kelvin, errors.New("not available")

	}
		
	//取得できない時は0 Kelvin
	log.Printf("openWeatherMap: %s: %.2f", city, d.Main.Kelvin)
	return d.Main.Kelvin, nil
}


type weatherUnderground struct {
	apiKey string
}

func (w weatherUnderground) temperature(city string) (float64, error) {
	urlstr := "http://api.wunderground.com/api/" + w.apiKey + "/conditions/q/" + city + ".json"
	//fmt.Println(urlstr)
	resp, err := http.Get(urlstr)
	if err != nil {
		//ここにはこない
		return 0, err
	}
	//apiKeyがなくてもここに来る
	defer resp.Body.Close()

	var d struct {
		Observation struct {
			Celsius float64 `json:"temp_c"`
		} `json:"current_observation"`
	}
	
	kelvin := 273.15
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		//デコードエラーになります。
		log.Printf("weatherUnderground: %s not available", city)
		return 0, err
	}
	kelvin = d.Observation.Celsius + 273.15
	log.Printf("weatherUnderground: %s: %.2f", city, kelvin)
	return kelvin, nil
}
