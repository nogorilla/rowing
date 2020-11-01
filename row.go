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

func calcLength(ts string) {
	t1, err := time.Parse("15:04:05", ts)
	t2, err := time.Parse("15:04:05", "00:00:00")

	fmt.Println(err, t1.Sub(t2))
	fmt.Println("hour", t1.Hour())
	fmt.Println("minute", t1.Minute())
	fmt.Println("second", t1.Second())
}

func calcPace(pace string) {
	paceTime, err := time.Parse("4:05.0", pace)
	zeroTime, err := time.Parse("4:05.0", "0:00.0")
	fmt.Println(err, paceTime.Sub(zeroTime))

	fmt.Println("hour", paceTime.Hour())
	fmt.Println("minute", paceTime.Minute())
	fmt.Println("second", paceTime.Second())
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
	var length string
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
				Usage:       "Distance travelled",
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
				Name:        "length",
				Usage:       "Time spent rowing",
				Aliases:     []string{"l"},
				Required:    true,
				Destination: &length,
			},
		},
		Action: func(c *cli.Context) error {
			// look for existing rowing.json file
			// if not create a new one

			// add entry to bottom of the file
			fmt.Printf("Distance: %x, Pace: %s, Date: %s\n", distance, pace, date)

			calcLength(length)
			calcPace(pace)

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
