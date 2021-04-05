package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v2"
)

func toSeconds(t string, format string, base string) float64 {
	duration, _ := time.Parse(format, t)
	zero, _ := time.Parse(format, base)

	result, _ := time.ParseDuration(duration.Sub(zero).String())

	return result.Seconds()
}

const (
	t1 = "01/02/2006"
	t2 = "1/02/2006"
	t3 = "01/2/2006"
	t4 = "1/1/2006"
)

func formatDate(d string) time.Time {
	var date time.Time
	var err error
	date, err = time.Parse(t1, d)
	if err != nil {
		date, err = time.Parse(t2, d)

		if err != nil {
			date, err = time.Parse(t3, d)

			if err != nil {
				date, err = time.Parse(t4, d)
			}
		}
	}

	return date
}

type Events struct {
	Events []Row `json:"events"`
}

type Row struct {
	Date     string  `json:"date"`
	Distance int     `json:"distance"`
	Duration float64 `json:"duration"`
	Actual   float64 `json:"actual"`
	Pace     float64 `json:"pace"`
	Power    int     `json:"power"`
	Notes    string  `json:"notes"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func bootstrap() *sql.DB {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	folder := usr.HomeDir
	file := ".rowing.db"
	database := folder + "/" + file
	needTable := false

	if _, err := os.Stat(database); os.IsNotExist(err) {
		f, err := os.Create(database) // Create SQLite file
		check(err)

		err = f.Close()
		check(err)

		fmt.Printf("file: %s created\n", database)
		needTable = true
	}

	db, _ := sql.Open("sqlite3", database)
	if needTable == true {
		createTable(db)
	}
	return db
}

func createTable(db *sql.DB) {
	createTableSql := `CREATE TABLE rowing (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"date" TEXT,
		"distance" INTEGER,
		"duration" INTEGER,
		"actual" INTEGER,
		"pace" INTEGER,
		"power" INTEGER
	);`

	log.Println("Create rowing table...")
	statement, err := db.Prepare(createTableSql) // Prepare SQL Statement
	check(err)

	_, err = statement.Exec()
	check(err)

	log.Println("rowing table created")
}

func writeEntry() {

}

func main() {
	var date string
	var distance int
	var power int
	var duration string
	var actual string
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
			&cli.StringFlag{
				Name:        "actual",
				Usage:       "Actual time spent rowing",
				Aliases:     []string{"a"},
				Required:    true,
				Destination: &actual,
			},
			&cli.IntFlag{
				Name:        "power",
				Usage:       "Power setting on rower",
				Aliases:     []string{"w"},
				Required:    false,
				Destination: &power,
				Value:       4,
			},
		},
		Action: func(c *cli.Context) error {

			var dSec float64
			var aSec float64
			var pSec float64

			db := bootstrap()

			// add entry to bottom of the file
			fmt.Printf("Distance: %x, Pace: %s, Date: %s\n", distance, pace, date)

			timeFmt := "04:05"
			zeroFmt := "00:00"

			if strings.Count(duration, ":") > 1 {
				timeFmt = "15:04:05"
				zeroFmt = "00:00:00"
			}

			dSec = toSeconds(duration, timeFmt, zeroFmt)
			if len(actual) > 0 {
				aSec = toSeconds(actual, timeFmt, zeroFmt)
			} else {
				aSec = dSec
			}

			// olive my love

			pSec = toSeconds(pace, "4:05.0", "0:00.0")

			fmt.Println("duration: lSec:", dSec)
			fmt.Println("duration: pSec:", pSec)
			fmt.Println("duration: aSec:", aSec)


			event := Row{
				Date:     date,
				Distance: distance,
				Duration: dSec,
				Actual:   aSec,
				Pace:     pSec,
				Power:    power,
				Notes:    notes,
			}


			fmt.Printf("date: %s, distance: %d, duration: %.0f, actual: %.0f, pace: %.1f\n", event.Date, event.Distance, event.Duration, event.Actual, event.Pace)


			insertSql := "INSERT INTO rowing(date, distance, duration, actual, pace, power) VALUES (?, ?, ?, ?, ?, ?)"
			stmt, err := db.Prepare(insertSql)
			check(err)

			_, err = stmt.Exec(event.Date, event.Distance, event.Duration, event.Actual, event.Pace, 4)
			check(err)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
