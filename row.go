package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/urfave/cli/v2"
)

// type Rowing struct {
// 	date     time.Time
// 	distance int64
// 	pace     int64
// }

func duration(ts string) {
	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)

	fmt.Println(t.Format("15:04:05.000"))

	h, m, s := t.Clock()
	ms := t.Nanosecond() / int(time.Millisecond)
	fmt.Printf("%02d:%02d:%02d.%03d\n", h, m, s, ms)
	fmt.Printf("%d", ms)
}

func fileExists() bool {
	userprofile := os.Getenv("USERPROFILE")
	info, err := os.Stat(path.Join(userprofile, "rowing.json"))
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {
	var date string
	var distance int64
	var pace string

	app := &cli.App{
		Name:  "rowing",
		Usage: "Tracking rowing exercises",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "date",
				Usage:       "Date of the exercise",
				Aliases:     []string{"t"},
				Value:       time.Now().UTC().Format("2006-01-02"),
				Destination: &date,
			},
			&cli.Int64Flag{
				Name:        "distance",
				Usage:       "Distance",
				Required:    true,
				Aliases:     []string{"d"},
				Destination: &distance,
			},
			&cli.StringFlag{
				Name:        "pace",
				Usage:       "Pace",
				Aliases:     []string{"p"},
				Required:    true,
				Destination: &pace,
			},
		},
		Action: func(c *cli.Context) error {
			// look for existing rowing.json file
			// if not create a new one

			// add entry to bottom of the file
			fmt.Printf("Distance: %x, Pace: %s, Date: %s\n", distance, pace, date)

			if fileExists() {
				fmt.Printf("file exist\n")
			} else {
				fmt.Printf("file doesn't exist exists\n")
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
