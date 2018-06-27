package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"./brute"
)

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
	flag.Parse()

	brute.Teste()
	fmt.Println("Teste: ", *url)
	fmt.Println("Teste: ", *endPoint)
	fmt.Println("teste: ", *cookies)

	validateInput(*url, *endPoint)
	validateURL(*url)
	validateURLendPoints(*url, *endPoint)

}
