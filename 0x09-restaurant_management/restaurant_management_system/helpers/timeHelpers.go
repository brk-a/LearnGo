package helpers

import (
    "time"
)

func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
	// if start.Before(end) {
	// 	return !check.Before(start) && !check.After(end)
	// }
	// if start.Equal(end) {
	// 	return check.Equal(start)
	// }
	// return !start.After(check) || !end.Before(check)
}