package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli"
)

var Version = "unknown"

type version struct {
	Major int
	Minor int
	Patch int
}

func (v version) increment(incrementField string) version {
	switch incrementField {
	case "Major":
		n := v.Major + 1
		return version{n, v.Minor, v.Patch}
	case "Minor":
		n := v.Minor + 1
		return version{v.Major, n, v.Patch}
	case "Patch":
		n := v.Patch + 1
		return version{v.Major, v.Minor, n}
	}
	return v
}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}

func getVersionLine(s []string) (int, string) {
	for i, v := range s {
		if match, _ := regexp.MatchString("version.*", v); match {
			return i, v
		}
	}
	return 0, "no dice"
}

func getVersion(s string) version {
	versionString := s
	versionArray := make([]int, 3)
	for i, v := range strings.Split(versionString, ".") {
		int, err := strconv.Atoi(v)
		check(err)
		versionArray[i] = int
	}
	version := version{
		versionArray[0],
		versionArray[1],
		versionArray[2],
	}
	return version
}

func bump(field string, version string) (string, error) {
	// Actual work
	if version == "" {
		return "", errors.New("No version string given. Check usage.")
	}
	versionNumber := getVersion(version)
	versionNumber = versionNumber.increment(field)
	return fmt.Sprintf("%v.%v.%v", versionNumber.Major, versionNumber.Minor, versionNumber.Patch), nil
}

func main() {
	// CLI Definition
	app := cli.NewApp()
	app.Name = "bump"
	app.Version = Version
	app.HelpName = "bump"
	app.Usage = "dumb version bump"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name: "noqqe",
		},
	}
	app.UsageText = "bump <command> version"
	app.Commands = []cli.Command{
		{
			Name:      "patch",
			ShortName: "p",
			Usage:     "increment the patch version",
			Action: func(c *cli.Context) error {
				data, err := bump("Patch", c.Args().Get(0))
				check(err)
				fmt.Println(data)
				return nil
			},
		},
		{
			Name:      "minor",
			ShortName: "m",
			Usage:     "increment the minor version",
			Action: func(c *cli.Context) error {
				data, err := bump("Minor", c.Args().Get(0))
				check(err)
				fmt.Println(data)
				return nil
			},
		},
		{
			Name:      "major",
			ShortName: "M",
			Usage:     "increment the major version",
			Action: func(c *cli.Context) error {
				data, err := bump("Major", c.Args().Get(0))
				check(err)
				fmt.Println(data)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
