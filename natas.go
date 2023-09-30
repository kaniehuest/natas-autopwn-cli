package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/urfave/cli/v2"
)

func level0() string {
	var username string = "natas0"
	var password string = "natas0"
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://natas0.natas.labs.overthewire.org", nil)
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

	re := regexp.MustCompile(`<!--The password for natas1 is (.*?) -->`)
	match := re.FindStringSubmatch(s)
	natas1Password := match[1]

	return natas1Password
}

func level1(password string) string {
	var username string = "natas1"
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://natas1.natas.labs.overthewire.org", nil)
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

	re := regexp.MustCompile(`<!--The password for natas2 is (.*?) -->`)
	match := re.FindStringSubmatch(s)
	natas2Password := match[1]

	return natas2Password
}

func level2(password string) string {
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

func main() {
	var level string

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "level",
				Aliases:     []string{"l"},
				Value:       "1",
				Usage:       "level to get password",
				Destination: &level,
			},
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.NArg() > 0 {
				level = cCtx.Args().Get(0)
			}

			if level == "0" {
				println("User: natas0")
				println("Password: natas0")
			} else if level == "1" {
				println("User: natas1")
				fmt.Printf("Password: %s\n", level0())
			} else if level == "2" {
				println("User: natas2")
				var natas1Password string = level0()
				fmt.Printf("Password: %s\n", level1(natas1Password))
			} else if level == "3" {
				println("User: natas3")
				var natas1Password string = level0()
				var natas2Password string = level1(natas1Password)
				fmt.Printf("Password: %s\n", level2(natas2Password))
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
