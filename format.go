package timex

import "time"

// Format
// layout DateTime: time.DateTime
// layout DateOnly: time.DateOnly
// layout TimeOnly: time.TimeOnly
func Format(t time.Time, layout string) string {
	return t.In(Loc()).Format(layout)
}
