// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"os/user"
// 	"path"
// 	"strconv"
// 	"strings"

// 	_ "github.com/mattn/go-sqlite3"
// )

// func main() {

// 	timeFmt := "04:05"
// 	zeroFmt := "00:00"

// 	usr, err := user.Current()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	jsonFile, _ := os.Open(path.Join(usr.HomeDir, "rowing.json"))
// 	defer jsonFile.Close()
// 	byteValue, _ := ioutil.ReadAll(jsonFile)
// 	var events []RowJson
// 	json.Unmarshal(byteValue, &events)

// 	db := bootstrap()

// 	for i := 0; i < len(events); i++ {
// 		if strings.Count(events[i].Duration, ":") > 1 {
// 			timeFmt = "15:04:05"
// 			zeroFmt = "00:00:00"
// 		}

// 		var pSec float64
// 		var aSec float64
// 		var dSec float64
// 		var power int

// 		dSec = toSeconds(events[i].Duration, timeFmt, zeroFmt)

// 		if len(events[i].Pace) > 0 {
// 			pSec = toSeconds(events[i].Pace, "4:05.0", "0:00.0")
// 		} else {
// 			pSec = 0.0
// 		}

// 		if len(events[i].Actual) > 0 {
// 			aSec = toSeconds(events[i].Actual, timeFmt, zeroFmt)
// 		} else {
// 			aSec = 0.0
// 		}

// 		power = events[i].Power

// 		date := formatDate(events[i].Date)
// 		// fmt.Println(date)
// 		distance, _ := strconv.Atoi(strings.Replace(strings.Replace(events[i].Distance, ",", "", -1), ".00", "", -1))
// 		notes := events[i].Notes

// 		// fmt.Println(pSec)
// 		// fmt.Printf("duration: %s, dSecs: %f, timefmt: %s, zeroFmt: %s\n", d, dSec, timeFmt, zeroFmt)
// 		// fmt.Printf("date: %s, distance: %d, length: %.0f, pace: %.1f\n", events[i].Date, distance, dSec, pSec)
// 		row := RowNew{date, distance, dSec, aSec, pSec, power, notes}
// 		fmt.Printf("date: %s, distance: %d, duration: %.0f, pace: %.1f\n", row.Date.Format("2006-01-02"), row.Distance, row.Duration, row.Pace)

// 		insertSql := "INSERT INTO rowing(date, distance, duration, actual, pace, power) VALUES (?, ?, ?, ?, ?, ?)"
// 		stmt, err := db.Prepare(insertSql)
// 		check(err)

// 		_, err = stmt.Exec(row.Date.Format("2006-01-02"), row.Distance, row.Duration, row.Duration, row.Pace, 4)
// 		check(err)

// 	}
// }
