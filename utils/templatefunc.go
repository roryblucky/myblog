package utils

import (
	"fmt"
	"time"
)

func DateFormat(time time.Time) string {
	s := time.Format("15:04:05")
	return fmt.Sprintf("%d年%d月%d日 %s", time.Year(), int(time.Month()), time.Day(), s)
}
