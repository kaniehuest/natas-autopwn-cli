package levels

import (
	b64 "encoding/base64"
	"encoding/hex"
	"html"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func getSecretFromLevel8(client *http.Client, username string, password string) string {
	req, err := http.NewRequest(http.MethodGet, "http://natas8.natas.labs.overthewire.org/index-source.html", nil)
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

	decodedHTML := html.UnescapeString(string(bodyText))

	re := regexp.MustCompile(`"([0-9a-fA-F]{32})"`)
	match := re.FindStringSubmatch(decodedHTML)
	secret := match[1]

	return secret
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func decodeSecretFromLevel8(secret string) string {
	decoded, err := hex.DecodeString(secret)
	if err != nil {
		log.Fatal(err)
	}

	reversed := reverseString(string(decoded))

	b64Decoded, err := b64.StdEncoding.DecodeString(reversed)
	if err != nil {
		log.Fatal(err)
	}

	return string(b64Decoded)
}

func GetLevel9Password(password string) string {
	var username string = "natas8"
	client := &http.Client{}

	var encodedSecret string = getSecretFromLevel8(client, username, password)
	var secret string = decodeSecretFromLevel8(encodedSecret)

	form := url.Values{}
	form.Add("secret", secret)
	form.Add("submit", "1")
	requestBody := strings.NewReader(form.Encode())

	req, err := http.NewRequest(http.MethodPost, "http://natas8.natas.labs.overthewire.org", requestBody)
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
	re := regexp.MustCompile(`The password for natas9 is (.*)`)
	match := re.FindStringSubmatch(s)
	natas9Passowrd := match[1]

	return natas9Passowrd
}
