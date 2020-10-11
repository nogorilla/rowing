package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

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
			fmt.Printf("Distance: %x, Pace: %s, Date: %s\n", distance, pace, date)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
