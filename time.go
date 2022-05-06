package time

import "time"

const (
	t1 = "01/02/2006"
	t2 = "1/02/2006"
	t3 = "01/2/2006"
	t4 = "1/1/2006"
)

func toSeconds(t string, format string, base string) float64 {
	duration, _ := time.Parse(format, t)
	zero, _ := time.Parse(format, base)

	result, _ := time.ParseDuration(duration.Sub(zero).String())

	return result.Seconds()
}

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
