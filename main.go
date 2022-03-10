package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	ua "github.com/wux1an/fake-useragent"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

const version = 1
const maxRequests = 5000
const APITargetList = "https://hutin-puy.nadom.app/hosts.json"

type FuckYouRussianShip struct {
	TargetList []string `json:"list"`
	Success    int      `json:"success"`
	Errors     int      `json:"errors"`
}

func (f *FuckYouRussianShip) GetTargetList() bool {
	req, err := http.NewRequest("GET", APITargetList, nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 1 * time.Second} // максимум 1 секунда ожидания ответа
	resp, err := client.Do(req)
	if err != nil {
		return false
	} else {
		defer resp.Body.Close()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	respString := string(body)

	switch resp.StatusCode {

	case 200:
		var result []string
		err := json.Unmarshal([]byte(respString), &result)
		if err != nil {
			return false
		}

		for _, item := range result {
			fmt.Println("", item)
		}

		break
	case 404:
		break
	default:
		break
	}

	return false
}

func (f *FuckYouRussianShip) GetTargetListWithProxy() {

	TARGET := "https://marzed.com"

	proxyAddress := "45.137.40.209:8762"

	proxyUrl, err := url.Parse("http://" + proxyAddress)
	//adding the proxy settings to the Transport object
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	auth := "perytmmv:72b7f5qclg3n"
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	transport.ProxyConnectHeader = http.Header{}
	transport.ProxyConnectHeader.Add("Proxy-Authorization", basicAuth)

	client := &http.Client{Transport: transport}

	////creating the proxyURL
	//proxyStr := "http://" + proxyAddress
	//proxyURL, err := url.Parse(proxyStr)
	//if err != nil {
	//	log.Println(err)
	//}
	//
	////creating the URL to be loaded through the proxy
	//urlStr := TARGET
	//urlTarget, err := url.Parse(urlStr)
	//if err != nil {
	//	log.Println(err)
	//}
	//
	////adding the proxy settings to the Transport object
	//transport := &http.Transport{
	//	Proxy: http.ProxyURL(proxyURL),
	//}
	//
	////adding the Transport object to the http Client
	//client := &http.Client{
	//	Transport: transport,
	//}

	//fmt.Println("client.Transport | ", client.Transport)

	//generating the HTTP GET request
	request, err := http.NewRequest("GET", TARGET, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("cf-visitor", "https")
	request.Header.Set("User-Agent", ua.Random())
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Accept-Language", "ru")
	request.Header.Set("x-forwarded-proto", "https")
	request.Header.Set("Accept-Encoding", "gzip, deflate, br")
	request.Header.Set("Upgrade-Insecure-Requests", "1")
	if err != nil {
		log.Println(err)
	}
	//request.SetBasicAuth("perytmmv", "72b7f5qclg3n")

	////adding proxy authentication
	//auth := proxyAuth
	//basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	//request.Header.Set("Proxy-Authorization", basicAuth)
	//

	//printing the request to the console
	dump, _ := httputil.DumpRequest(request, false)
	fmt.Println("dump:", string(dump))

	//calling the URL
	response, err := client.Do(request)
	if err != nil {
		log.Println("[ERROR]", err)
		return
	}

	log.Println(response.StatusCode)
	log.Println(response.Status)
	//getting the response
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	//printing the response
	log.Println(string(data))
}

func main() {
	fmt.Println("Hello")
	var Kurwa FuckYouRussianShip
	Kurwa.GetTargetListWithProxy()
}
