// package

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

// 	var eventsNew []RowNew

// 	for i := 0; i < len(events); i++ {
// 		date := formatDate(events[i].Date)
// 		// fmt.Println(date)
// 		distance, _ := strconv.Atoi(strings.Replace(strings.Replace(events[i].Distance, ",", "", -1), ".00", "", -1))
// 		d := events[i].Duration
// 		p := events[i].Pace
// 		power := events[i].Power
// 		notes := events[i].Notes

// 		if strings.Count(d, ":") > 1 {
// 			timeFmt = "15:04:05"
// 			zeroFmt = "00:00:00"
// 		}

// 		var pSec float64

// 		dSec := toSeconds(d, timeFmt, zeroFmt)
// 		if len(p) > 0 {
// 			pSec = toSeconds(p, "4:05.0", "0:00.0")
// 		} else {
// 			pSec = 0.0
// 		}
// 		// fmt.Println(pSec)
// 		// fmt.Printf("duration: %s, dSecs: %f, timefmt: %s, zeroFmt: %s\n", d, dSec, timeFmt, zeroFmt)
// 		// fmt.Printf("date: %s, distance: %d, length: %.0f, pace: %.1f\n", events[i].Date, distance, dSec, pSec)
// 		r := RowNew{date, distance, dSec, pSec, power, notes}
// 		eventsNew = append(eventsNew, r)
// 	}

// 	for i := 0; i < len(events); i++ {
// 		fmt.Println(events[i])
// 		//fmt.Printf("date: %s, distance: %d, duration: %.0f, pace: %.1f\n", events[i].Date, events[i].Distance, events[i].Duration, events[i].Pace)
// 	}

// 	file, _ := json.MarshalIndent(eventsNew, "", " ")
// 	_ = ioutil.WriteFile(path.Join(usr.HomeDir, "rowing-new.json"), file, 0644)
// }
