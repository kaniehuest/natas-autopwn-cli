package levels

import (
	"io"
	"log"
	"net/http"
	"regexp"
)

func GetLevel1Password() string {
	var username string = "natas0"
	var password string = "natas0"
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://natas0.natas.labs.overthewire.org", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	s := string(bodyText)

	re := regexp.MustCompile(`<!--The password for natas1 is (.*?) -->`)
	match := re.FindStringSubmatch(s)
	natas1Password := match[1]

	return natas1Password
}
