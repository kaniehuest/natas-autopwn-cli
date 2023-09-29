package main

import (
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
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := io.ReadAll(resp.Body)
	s := string(bodyText)

	re := regexp.MustCompile(`<!--The password for natas1 is (.*?) -->`)
	match := re.FindStringSubmatch(s)
	natas1Password := match[1]

	return natas1Password
}

func level1() string {
	var username string = "natas1"
	var password string = level0()
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://natas1.natas.labs.overthewire.org", nil)
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := io.ReadAll(resp.Body)
	s := string(bodyText)

	re := regexp.MustCompile(`<!--The password for natas2 is (.*?) -->`)
	match := re.FindStringSubmatch(s)
	natas2Password := match[1]

	return natas2Password
}

func main() {
	var level string

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "level",
				Aliases:     []string{"l"},
				Value:       "0",
				Usage:       "level to get password",
				Destination: &level,
			},
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.NArg() > 0 {
				level = cCtx.Args().Get(0)
			}

			if level == "0" {
				println(level0())
			} else if level == "1" {
				println(level1())
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
