package util

import (
	"fmt"
	"strconv"
	"time"
)

func CopyrightYear(since string) string {
	nowYear := strconv.Itoa(time.Now().Year())
	if since == "" || nowYear == since {
		return nowYear
	} else {
		return fmt.Sprintf("%s - %s", since, nowYear)
	}
}
