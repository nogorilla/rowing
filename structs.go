// package structs

// import "time"

// type Row struct {
// 	Date     string  `json:"date"`
// 	Distance int     `json:"distance"`
// 	Duration float64 `json:"duration"`
// 	Actual   float64 `json:"actual"`
// 	Pace     float64 `json:"pace"`
// 	Power    int     `json:"power"`
// 	Notes    string  `json:"notes"`
// }

// type RowNew struct {
// 	Date     time.Time `json:"date"`
// 	Distance int       `json:"distance"`
// 	Duration float64   `json:"duration"`
// 	Actual   float64   `json:"actual"`
// 	Pace     float64   `json:"pace"`
// 	Power    int       `json:"power"`
// 	Notes    string    `json:"notes"`
// }

// type RowJson struct {
// 	Date     string `json:"date"`
// 	Distance string `json:"distance"`
// 	Duration string `json:"duration"`
// 	Actual   string `json:"actual"`
// 	Pace     string `json:"pace"`
// 	Power    int    `json:"power"`
// 	Notes    string `json:"notes"`
// }

// type Events struct {
// 	Events []Row `json:"events"`
// }
