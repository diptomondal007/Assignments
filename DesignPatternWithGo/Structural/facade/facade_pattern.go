package facade

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	commonRequestPrefix              = "http://api.openweathermap.org/data/2.5/"
	weatherByCityName                = commonRequestPrefix + "weather?q=%s,%s&APPID=%s"
	weatherByGeographicalCoordinates = commonRequestPrefix + "weather?lat=%f&lon=%f&APPID=%s"
)

type CurrentWeatherDataRetriever interface {
	GetByCityAndCountryCode(city, countryCode string) (Weather, error)
	GetByGeoCoordinates(lat, lon float32) (Weather, error)
}

type Weather struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Timezone int    `json:"timezone"`
	Cod      int    `json:"cod"`
	Coord    struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
}

type CurrentWeatherData struct {
	APIkey string
}


func (p *CurrentWeatherData) responseParser(body io.Reader) (*Weather, error) {
	w := new(Weather)
	err := json.NewDecoder(body).Decode(w)
	if err != nil{
		return nil, err
	}
	return w, nil
}

//GetByGeoCoordinates returns the current weather data by passing a geographical
//coordinates (latitude and longitude) in decimal notation. It returns weather
//information or a detailed error.
//For example, to query about Madrid, Spain you do:
//	currentWeather.GetByGeoCoordinates(-3, 40)
func (c *CurrentWeatherData) GetByGeoCoordinates(lat, lon float32) (weather *Weather, err error) {
	return c.doRequest(fmt.Sprintf(weatherByGeographicalCoordinates, lat, lon, c.APIkey))
}

//GetByCityAndCountryCode returns the current weather data by passing a city name
//and an ISO country code. It returns weather information or a detailed error
//For example, to query about Madrid, Spain you do:
//	currentWeather.GetByCityAndCountryCode("Madrid", "ES)
func (c *CurrentWeatherData) GetByCityAndCountryCode(city, countryCode string) (weather *Weather, err error) {
	return c.doRequest(fmt.Sprintf(weatherByCityName, city, countryCode, c.APIkey))
}

func (o *CurrentWeatherData) doRequest(uri string) (weather *Weather, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		byt, errMsg := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if errMsg == nil {
			errMsg = fmt.Errorf("%s", string(byt))
		}
		err = fmt.Errorf("Status code was %d, aborting. Error message was:\n%s\n",
			resp.StatusCode, errMsg)

		return
	}

	weather, err = o.responseParser(resp.Body)
	resp.Body.Close()

	return
}