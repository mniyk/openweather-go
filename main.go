package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/go-ini/ini"
)

const END_POINT = "https://api.openweathermap.org/data/2.5/weather"

type Response struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int64     `josn:"timezone"`
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Cod        int64     `json:"cod"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	Id          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
	SeaLevel  int64   `json:"sea_level"`
	GrndLevel int64   `json:"grnd_level"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Sys struct {
	Type    int64  `json:"type"`
	Id      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

func main() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Failed to read config file %s.", err)
	}

	token := cfg.Section("OPEN_WEATHER").Key("API_KEY").String()

	city := "Osaka,jp"

	values := url.Values{}
	values.Set("q", city)
	values.Set("APPID", token)

	url := END_POINT + "?" + values.Encode()

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var data Response
	if err := json.Unmarshal(bytes, &data); err != nil {
		panic(err)
	}

	fmt.Printf("%+v", data)
}
