package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

func toSeconds(t string, format string, base string) float64 {
	duration, _ := time.Parse(format, t)
	zero, _ := time.Parse(format, base)

	result, _ := time.ParseDuration(duration.Sub(zero).String())

	return result.Seconds()
}

func fileExists() bool {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	info, err := os.Stat(path.Join(usr.HomeDir, "rowing.json"))
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

type Events struct {
	Events []Row `json:"events"`
}

type Row struct {
	Date     string  `json:"date"`
	Distance int     `json:"distance"`
	Duration float64 `json:"duration"`
	Pace     float64 `json:"pace"`
	Power    int     `json:"power"`
	Notes    string  `json:"notes"`
}

func main() {
	var date string
	var distance int
	var power int
	var duration string
	var pace string
	power, notes := 4, ""

	app := &cli.App{
		Name:  "rowing",
		Usage: "Tracking rowing exercises",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "date",
				Usage:       "Date of the exercise",
				Aliases:     []string{"t"},
				Value:       time.Now().UTC().Format("2006-01-02T00:00:00Z"),
				Destination: &date,
			},
			&cli.IntFlag{
				Name:        "distance",
				Usage:       "Distance traveled",
				Required:    true,
				Aliases:     []string{"d"},
				Destination: &distance,
			},
			&cli.StringFlag{
				Name:        "pace",
				Usage:       "Pace per 500 meters",
				Aliases:     []string{"p"},
				Required:    true,
				Destination: &pace,
			},
			&cli.StringFlag{
				Name:        "duration",
				Usage:       "Time spent rowing",
				Aliases:     []string{"l"},
				Required:    true,
				Destination: &duration,
			},
			&cli.IntFlag{
				Name:        "power",
				Usage:       "Power setting on rower",
				Aliases:     []string{"w"},
				Required:    false,
				Destination: &power,
				Value:       4,
			},
			&cli.StringFlag{
				Name:        "notes",
				Usage:       "Notes for rowing event",
				Aliases:     []string{"n"},
				Required:    false,
				Destination: &notes,
				Value:       "",
			},
		},
		Action: func(c *cli.Context) error {
			// add entry to bottom of the file
			fmt.Printf("Distance: %x, Pace: %s, Date: %s\n", distance, pace, date)

			timeFmt := "04:05"
			zeroFmt := "00:00"

			if strings.Count(duration, ":") > 1 {
				timeFmt = "15:04:05"
				zeroFmt = "00:00:00"
			}

			dSec := toSeconds(duration, timeFmt, zeroFmt)
			pSec := toSeconds(pace, "4:05.0", "0:00.0")

			fmt.Println("duration: lSec:", dSec)
			fmt.Println("duration: pSec:", pSec)

			usr, err := user.Current()
			if err != nil {
				log.Fatal(err)
			}

			jsonFile, _ := os.Open(path.Join(usr.HomeDir, "rowing-new.json"))
			defer jsonFile.Close()
			byteValue, _ := ioutil.ReadAll(jsonFile)

			var events []Row
			json.Unmarshal(byteValue, &events)

			event := Row{
				Date:     date,
				Distance: distance,
				Duration: dSec,
				Pace:     pSec,
				Power:    power,
				Notes:    notes,
			}

			events = append(events, event)

			// fmt.Println(event)
			// t, _ := time.Parse("2006-01-02", date)
			// fmt.Println(t)
			// for i := 0; i < len(events); i++ {
			// 	fmt.Printf("date: %s, distance: %d, duration: %.0f, pace: %.1f\n", events[i].Date, events[i].Distance, events[i].Duration, events[i].Pace)
			// }

			file, _ := json.MarshalIndent(events, "", " ")
			_ = ioutil.WriteFile(path.Join(usr.HomeDir, "rowing-new.json"), file, 0644)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
