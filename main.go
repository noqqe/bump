package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli"
)

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
		fmt.Println(e)
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
	versionString := strings.Split(s, "'")[1]
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

func bump(field string, file string) (string, error) {
	// Actual work
	dat, err := ioutil.ReadFile(file)
	check(err)
	metadata := strings.Split(string(dat), "\n")
	line, version := getVersionLine(metadata)
	versionNumber := getVersion(version)
	versionNumber = versionNumber.increment(field)
	metadata[line] = fmt.Sprintf("version '%v.%v.%v'", versionNumber.Major, versionNumber.Minor, versionNumber.Patch)
	return strings.Join(metadata[:len(metadata)-1], "\n"), nil
}

func main() {
	// CLI Definition
	app := cli.NewApp()
	app.Name = "Incrementer"
	app.Version = "0.1.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Jason Morgan",
			Email: "Jason.Morgan@digitalglobe.com",
		},
	}
	app.UsageText = "incrementer COMMAND PATH_TO_METADATA_FILE"
	app.Commands = []cli.Command{
		{
			Name:      "patch",
			ShortName: "p",
			Usage:     "increment the patch version",
			Action: func(c *cli.Context) error {
				data, err := bump("Patch", os.Args[2])
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
				data, err := bump("Minor", os.Args[2])
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
				data, err := bump("Major", os.Args[2])
				check(err)
				fmt.Println(data)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
