package util

import (
	"fmt"
	"strconv"
	"time"
)

func CopyrightYear(since int) string {
	nowYear := time.Now().Year()
	if since == 0 || nowYear == since {
		return strconv.Itoa(nowYear)
	} else {
		return fmt.Sprintf("%d - %d", since, nowYear)
	}
}
