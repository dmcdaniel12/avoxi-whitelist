package main

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/azer/logger"
	"github.com/oschwald/maxminddb-golang"
)

type Whitelist struct {
	Ip        string   `json: "ip"`
	Countries []string `json: "countries`
}

type Response struct {
	Status   string `json: "status"`
	Response string `json: "response`
}

var Record struct {
	Country struct {
		ISOCode string `maxminddb:"iso_code"`
	} `maxminddb:"country"`
}

func CheckIsWhitelisted(w http.ResponseWriter, r *http.Request) {
	var log = logger.New("avoxi-whitelist")
	log.Info("CheckIsWhitelisted")

	reqBody, _ := ioutil.ReadAll(r.Body)

	var whitelist Whitelist
	json.Unmarshal(reqBody, &whitelist)

	// Check struct to see if it matches maxmind ip addresses
	db, err := maxminddb.Open("maxmind/GeoLite2-Country.mmdb")

	if err != nil {
		log.Error(err.Error())
	}

	defer db.Close()

	ip := net.ParseIP(whitelist.Ip)

	err = db.Lookup(ip, &Record)

	if err != nil {
		log.Error(err.Error())
	}

	result := countryExists(Record.Country.ISOCode, whitelist.Countries)

	var response Response

	if result {
		response.Status = "Success"
		response.Response = "Found the country"
	} else {
		response.Status = "Failed"
		response.Response = "Country not found in the list"
	}

	json.NewEncoder(w).Encode(response)
}

func countryExists(country string, countries []string) (result bool) {
	result = false
	for _, c := range countries {
		if c == country {
			result = true
			break
		}
	}
	return result
}
