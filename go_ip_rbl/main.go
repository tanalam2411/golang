package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
https://www.upwork.com/jobs/~01ca4dc7158efed30a
Mock API Get Ip list - https://www.mockable.io/a/#/space/demo7465998/rest/UAAAAAAAA?inwizzard=true
Mock API Get proxy list - https://www.mockable.io/a/#/space/demo7465998/rest/UAAAAAAAA?inwizzard=true

passwd - rbl0011check
*/

/*
getIPurl ->
	HTTP method : GET
	Request body: {}
	Response : {
					"ip": ["x.x.x.x", "y.y.y.y", ...],  // list of ip
					"ip_range": ["172.16.2.x-172.16.2.y", "172.17.3.x-172.16.3.y", ...], // Provide ip range to generte ip list
					"cidr": ["10.1.1.5/24", ...],
					.
					.
					.
				}

getPROXYurl ->
	HTTP method : GET
	Request body : {}
	Response : {
					[{"url": "/api/v1/ip/blacklisted", "http_method_type": "get/post", "params": {}}]
				}

*/

var (
	defaultGetIPURL   = "https://demo7465998.mockable.io/ips"
	defaultProxyURL   = "https://demo7465998.mockable.io/proxy/urls"
	defaultWebhookURL = "https://demo7465998.mockable.io/webhook"
)

func getIPList(url string) []byte {

	resp, err := http.Get(url)

	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	ipList, err := ioutil.ReadAll(resp.Body)

	// logic to create ip_list based on provided params (ip_range, cird etc)

	// return ipList
	return ipList
}

func checkRBL(ip string, getProxyURL string) bool {
	// steps:
	// 1) Get all proxy urls
	// 2) Iterate over it and check RBL status of IP using proxy url, if got response then break
	// 3) Return RBL status
}

func updateStatus(ip string, status bool, webhookURL string) bool {

	// http call to webhookurl to update ip's status
	// return update status
}

func main() {

	getIPURL := flag.String("ipUrl", defaultGetIPURL, "URL to fetch ip list")

	getProxyURL := flag.String("proxyUrl", defaultProxyURL, "URL to fetch proxy url list")

	webhookURL := flag.String("proxyUrl", defaultWebhookURL, "Webhook url to update status of ip")

	flag.Parse()

	ipList := getIPList(*getIPURL)

	for _, ip := range ipList {

		status = checkRBL(ip, getProxyURL)

		updateStatus(ip, status, webhookURL)
	}

}
