package levels

import (
	"io"
	"log"
	"net/http"
	"regexp"
)

func GetLevel3Password(password string) string {
	var username string = "natas2"
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://natas2.natas.labs.overthewire.org/files/users.txt", nil)
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

	re := regexp.MustCompile(`natas3:(.*)`)
	match := re.FindStringSubmatch(s)
	natas3Password := match[1]

	return natas3Password
}
