package glob

import "time"

func GetNowUTC() (todayUtc int64) {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	sec := now.Second()

	today := time.Date(year, month, day, hour, minute, sec, 0, time.Local)
	todayUtc = today.UTC().Unix()
	return
}
