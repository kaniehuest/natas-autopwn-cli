package levels

import (
	"io"
	"log"
	"net/http"
	"regexp"
)

func GetLevel4Password(password string) string {
	var username string = "natas3"
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://natas3.natas.labs.overthewire.org/s3cr3t/users.txt", nil)
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

	re := regexp.MustCompile(`natas4:(.*)`)
	match := re.FindStringSubmatch(s)
	natas4Password := match[1]

	return natas4Password
}
