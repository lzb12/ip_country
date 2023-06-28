package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type JSONData struct {
	IP string `json:"ip"`
	Network string `json:"network"`
	Version string `json:"version"`
	City string `json:"city"`
	Region string `json:"region"`
	RegionCode string `json:"region_code"`
	Country string `json:"country"`
	CountryName string `json:"country_name"`
	CountryCode string `json:"country_code"`
	CountryCodeIso3 string `json:"country_code_iso3"`
	CountryCapital string `json:"country_capital"`
	CountryTld string `json:"country_tld"`
	ContinentCode string `json:"continent_code"`
	InEu bool `json:"in_eu"`
	Postal interface{} `json:"postal"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone string `json:"timezone"`
	UtcOffset string `json:"utc_offset"`
	CountryCallingCode string `json:"country_calling_code"`
	Currency string `json:"currency"`
	CurrencyName string `json:"currency_name"`
	Languages string `json:"languages"`
	CountryArea float64 `json:"country_area"`
	CountryPopulation int `json:"country_population"`
	Asn string `json:"asn"`
	Org string `json:"org"`
}

func main() {
	filePath := "ip_addresses.txt"

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ip := scanner.Text()
		// 构建请求URL
		url := fmt.Sprintf("https://ipapi.co/%s/json/", ip)
		//fmt.Println(url)
		// 发起GET请求
		response, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		// 读取响应内容
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		// 解析JSON数据
		var ipInfo JSONData
		err = json.Unmarshal(body, &ipInfo)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(ipInfo)
		// 获取城市信息
		country := ipInfo.Country
		time.Sleep(time.Minute)

		// 如果城市不等于"Sydney"，则打印城市信息
		if country == "AU" {
			fmt.Printf("%s, country: %s\n", ip, country)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
