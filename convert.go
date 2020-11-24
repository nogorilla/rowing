package main

import (
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
	Pace     string `json:"pace"`
	Power    int    `json:"power"`
	Notes    string `json:"notes"`
}

type RowNew struct {
	Date     string  `json:"date"`
	Distance int     `json:"distance"`
	Duration float64 `json:"duration"`
	Pace     float64 `json:"pace"`
	Power    int     `json:"power"`
	Notes    string  `json:"notes"`
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

	var eventsNew []RowNew

	for i := 0; i < len(events); i++ {
		date := events[i].Date
		distance, _ := strconv.Atoi(strings.Replace(events[i].Distance, ",", "", -1))
		d := events[i].Duration
		p := events[i].Pace
		power := events[i].Power
		notes := events[i].Notes

		if strings.Count(d, ":") > 1 {
			timeFmt = "15:04:05"
			zeroFmt = "00:00:00"
		}

		var pSec float64

		dSec := toSecs(d, timeFmt, zeroFmt)
		if len(p) > 0 {
			pSec = toSecs(p, "4:05.0", "0:00.0")
		} else {
			pSec = 0.0
		}
		fmt.Println(pSec)
		// fmt.Printf("duration: %s, dSecs: %f, timefmt: %s, zeroFmt: %s\n", d, dSec, timeFmt, zeroFmt)
		// fmt.Printf("date: %s, distance: %d, length: %.0f, pace: %.1f\n", events[i].Date, distance, dSec, pSec)
		r := RowNew{date, distance, dSec, pSec, power, notes}
		eventsNew = append(eventsNew, r)
	}

	fmt.Println(eventsNew)

	file, _ := json.MarshalIndent(eventsNew, "", " ")
	_ = ioutil.WriteFile(path.Join(usr.HomeDir, "rowing-new.json"), file, 0644)
}
