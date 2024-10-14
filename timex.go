package timex

import (
	"time"
)

func Loc() *time.Location {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return loc
}

func Now() time.Time {
	return time.Now().In(Loc())
}

func NowUnix() int64 {
	return time.Now().In(Loc()).Unix()
}

// ThisWeekRange returns the start and end time of this week
func ThisWeekRange() (start time.Time, end time.Time) {
	loc := Loc()
	now := time.Now().In(loc)
	weekday := now.Weekday()

	daysSinceSunday := int(weekday)

	// get last sunday date
	sundayDate := now.AddDate(0, 0, -daysSinceSunday)
	year, month, day := sundayDate.Date()

	start = time.Date(year, month, day, 0, 0, 0, 0, loc)

	// end date is start + 7 days - 1 second
	end = start.Add(7*24*time.Hour - time.Second)
	return start, end
}

func TodayRange() (start time.Time, end time.Time) {
	loc := Loc()
	now := time.Now().In(loc)
	start = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	return start, start.Add(24*time.Hour - time.Second)
}

func Unix(sec int64, nsec int64) time.Time {
	return time.Unix(sec, nsec).In(Loc())
}

// FormatUnix
// layout DateTime: time.DateTime
// layout DateOnly: time.DateOnly
// layout TimeOnly: time.TimeOnly
func FormatUnix(sec int64, nsec int64, layout string) string {
	return Unix(sec, nsec).Format(layout)
}
