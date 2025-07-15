package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v3"
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
		return version{n, 0, 0}
	case "Minor":
		n := v.Minor + 1
		return version{v.Major, n, 0}
	case "Patch":
		n := v.Patch + 1
		return version{v.Major, v.Minor, n}
	}
	return v
}

func getVersion(s string) (version, error) {
	versionString := s
	versionArray := make([]int, 3)
	for i, v := range strings.Split(versionString, ".") {
		int, err := strconv.Atoi(v)

		// Check if we have a valid version string
		// if it could not beconverted its probably a string
		if err != nil {
			return version{}, errors.New("invalid version string: " + s)
		}
		versionArray[i] = int
	}
	version := version{
		versionArray[0],
		versionArray[1],
		versionArray[2],
	}
	return version, nil
}

func bump(field string, version string) (string, error) {
	versionNumber, err := getVersion(version)
	if err != nil {
		return version, nil
	}
	versionNumber = versionNumber.increment(field)
	return fmt.Sprintf("%v.%v.%v", versionNumber.Major, versionNumber.Minor, versionNumber.Patch), nil
}

func main() {
	// CLI Definition
	cmd := &cli.Command{
		Name:                  "bump",
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "patch",
				Aliases: []string{"p"},
				Usage:   "increment the patch version",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					data, err := bump("Patch", cmd.Args().Get(0))
					if err != nil {
						return cli.Exit("Could not bump version", 1)
					}
					fmt.Println(data)
					return nil
				},
			},
			{
				Name:    "minor",
				Aliases: []string{"m"},
				Usage:   "increment the minor version",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					data, err := bump("Minor", cmd.Args().Get(0))
					if err != nil {
						return cli.Exit("Could not bump version", 1)
					}
					fmt.Println(data)
					return nil
				},
			},
			{
				Name:    "major",
				Aliases: []string{"M"},
				Usage:   "increment the major version",
				Action: func(c context.Context, cmd *cli.Command) error {
					data, err := bump("Major", cmd.Args().Get(0))
					if err != nil {
						return cli.Exit("Could not bump version", 1)
					}
					fmt.Println(data)
					return nil
				},
			},
		},
	}
	cmd.Run(context.Background(), os.Args)
}
