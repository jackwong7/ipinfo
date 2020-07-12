package ipinfo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type IpJson struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Timezone string `json:"timezone"`
}

func GetIp() IpJson {
	defer func() {
		recover()
	}()
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	req, _ := http.NewRequest("GET", "https://ipinfo.io", nil)
	req.Header.Set("Authorization", "Bearer 42b543091592e3")
	resp, err := client.Do(req)

	if err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			ipJsonObj := IpJson{}
			json.Unmarshal(body, &ipJsonObj)
			return ipJsonObj
		}
	}
	return IpJson{}
}
