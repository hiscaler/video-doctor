package datetime

import "fmt"

func Timedelta(seconds int) string {
	minutes := seconds / 60
	seconds = seconds % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
