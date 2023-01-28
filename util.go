package tetris

import (
	"time"
)

// sleep_milli sleeps for the specified number of milliseconds
func sleep_milli(millis int) {
	n := time.Duration(millis) * time.Millisecond
	time.Sleep(n)
}
