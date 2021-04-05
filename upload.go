package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func toSecs(t string, format string, base string) float64 {
	duration, _ := time.Parse(format, t)
	zero, _ := time.Parse(format, base)

	result, _ := time.ParseDuration(duration.Sub(zero).String())

	return result.Seconds()
}

type Events struct {
	Events []Row `json:"events"`
}

type Row struct {
	Date     string `json:"date"`
	Distance string `json:"distance"`
	Duration string `json:"duration"`
	Actual   string `json:"actual"`
	Pace     string `json:"pace"`
	Power    int    `json:"power"`
	Notes    string `json:"notes"`
}

type RowNew struct {
	Date     time.Time `json:"date"`
	Distance int       `json:"distance"`
	Duration float64   `json:"duration"`
	Actual   float64   `json:"actual"`
	Pace     float64   `json:"pace"`
	Power    int       `json:"power"`
	Notes    string    `json:"notes"`
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

func main() {

	timeFmt := "04:05"
	zeroFmt := "00:00"

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	jsonFile, _ := os.Open(path.Join(usr.HomeDir, "rowing.json"))
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var events []Row
	json.Unmarshal(byteValue, &events)

	db := bootstrap()

	for i := 0; i < len(events); i++ {
		if strings.Count(events[i].Duration, ":") > 1 {
			timeFmt = "15:04:05"
			zeroFmt = "00:00:00"
		}

		var pSec float64
		var aSec float64
		var dSec float64
		var power int

		dSec = toSecs(events[i].Duration, timeFmt, zeroFmt)

		if len(events[i].Pace) > 0 {
			pSec = toSecs(events[i].Pace, "4:05.0", "0:00.0")
		} else {
			pSec = 0.0
		}

		if len(events[i].Actual) > 0 {
			aSec = toSecs(events[i].Actual, timeFmt, zeroFmt)
		} else {
			aSec = 0.0
		}

		power = events[i].Power

		date := formatDate(events[i].Date)
		// fmt.Println(date)
		distance, _ := strconv.Atoi(strings.Replace(strings.Replace(events[i].Distance, ",", "", -1), ".00", "", -1))
		notes := events[i].Notes

		// fmt.Println(pSec)
		// fmt.Printf("duration: %s, dSecs: %f, timefmt: %s, zeroFmt: %s\n", d, dSec, timeFmt, zeroFmt)
		// fmt.Printf("date: %s, distance: %d, length: %.0f, pace: %.1f\n", events[i].Date, distance, dSec, pSec)
		row := RowNew{date, distance, dSec, aSec, pSec, power, notes}
		fmt.Printf("date: %s, distance: %d, duration: %.0f, pace: %.1f\n", row.Date.Format("2006-01-02"), row.Distance, row.Duration, row.Pace)

		insertSql := "INSERT INTO rowing(date, distance, duration, actual, pace, power) VALUES (?, ?, ?, ?, ?, ?)"
		stmt, err := db.Prepare(insertSql)
		check(err)

		_, err = stmt.Exec(row.Date.Format("2006-01-02"), row.Distance, row.Duration, row.Duration, row.Pace, 4)
		check(err)

	}
}
