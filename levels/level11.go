package levels

import (
	b64 "encoding/base64"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

func makeValidCookieForLevel11(XORKey string) string {
	const newCookie string = `{"showpassword":"yes","bgcolor":"#ffffff"}`

	encodedCookie := make([]byte, len(newCookie))
	for i := 0; i < len(newCookie); i++ {
		encodedCookie[i] += newCookie[i] ^ XORKey[i%4]
	}

	b64EncodedCookie := b64.StdEncoding.EncodeToString(encodedCookie)

	return b64EncodedCookie
}

func getLevel11XORKey(b64EncodedCookie string) string {
	b64DecodedCookie, err := b64.StdEncoding.DecodeString(b64EncodedCookie)
	if err != nil {
		log.Fatal(err)
	}

	const cookieArray string = `{"showpassword":"no","bgcolor":"#ffffff"}`

	key := make([]byte, len(cookieArray))
	for i := 0; i < len(cookieArray); i++ {
		key[i] = cookieArray[i] ^ b64DecodedCookie[i]
	}

	decodedKey := string(key[:4])

	return decodedKey
}

func getLevel11Cookie(client *http.Client, username string, password string) string {
	req, err := http.NewRequest(http.MethodPost, "http://natas11.natas.labs.overthewire.org", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	cookies := resp.Cookies()

	encodedCookie := string(cookies[0].Value)
	encodedCookie, err = url.QueryUnescape(encodedCookie)
	if err != nil {
		log.Fatal(err)
	}

	return encodedCookie
}

func GetLevel12Password(password string) string {
	var username string = "natas11"
	client := &http.Client{}

	var level11Cookie string = getLevel11Cookie(client, username, password)
	var XORKey string = getLevel11XORKey(level11Cookie)
	var validCookie string = makeValidCookieForLevel11(XORKey)

	req, err := http.NewRequest(http.MethodPost, "http://natas11.natas.labs.overthewire.org", nil)
	if err != nil {
		log.Fatal(err)
	}

	cookie := &http.Cookie{
		Name:  "data",
		Value: validCookie,
	}
	req.AddCookie(cookie)
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
	re := regexp.MustCompile(`The password for natas12 is (.*)<br>`)
	match := re.FindStringSubmatch(s)
	natas12Password := match[1]

	return natas12Password
}
