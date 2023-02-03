package tetris

import (
	"time"
)

// Sleep sleeps for the specified number of milliseconds
func Sleep(millis int) {
	n := time.Duration(millis) * time.Millisecond
	time.Sleep(n)
}
