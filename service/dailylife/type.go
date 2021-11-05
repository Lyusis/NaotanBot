package dailylife

import "time"

var (
	weekdayCN = [7]string{"日", "一", "二", "三", "四", "五", "六"}
)

type weathers struct {
	Results []struct {
		Location struct {
			Id             string `json:"id"`
			Name           string `json:"name"`
			Country        string `json:"country"`
			Path           string `json:"path"`
			Timezone       string `json:"timezone"`
			TimezoneOffset string `json:"timezone_offset"`
		} `json:"location"`
		Daily []struct {
			Date                string `json:"date"`
			TextDay             string `json:"text_day"`
			CodeDay             string `json:"code_day"`
			TextNight           string `json:"text_night"`
			CodeNight           string `json:"code_night"`
			High                string `json:"high"`
			Low                 string `json:"low"`
			Rainfall            string `json:"rainfall"`
			Precip              string `json:"precip"`
			WindDirection       string `json:"wind_direction"`
			WindDirectionDegree string `json:"wind_direction_degree"`
			WindSpeed           string `json:"wind_speed"`
			WindScale           string `json:"wind_scale"`
			Humidity            string `json:"humidity"`
		} `json:"daily"`
		LastUpdate time.Time `json:"last_update"`
	} `json:"results"`
}

type news struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Newslist []struct {
		Mtime  string `json:"mtime"`
		Title  string `json:"title"`
		Digest string `json:"digest"`
		Imgsrc string `json:"imgsrc"`
		Url    string `json:"url"`
		Source string `json:"source"`
	} `json:"newslist"`
}
