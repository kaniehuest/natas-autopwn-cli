package levels

import (
	"io"
	"log"
	"net/http"
	"regexp"
)

func GetLevel2Password(password string) string {
	var username string = "natas1"
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://natas1.natas.labs.overthewire.org", nil)
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

	re := regexp.MustCompile(`<!--The password for natas2 is (.*?) -->`)
	match := re.FindStringSubmatch(s)
	natas2Password := match[1]

	return natas2Password
}
