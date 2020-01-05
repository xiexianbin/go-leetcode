package t1154

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func dayOfYear(date string) int {
	mDays := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}
	dL := strings.Split(date, "-")
	y, err := strconv.Atoi(dL[0])
	if err != nil {
		return -1
	}
	m, err := strconv.Atoi(dL[1])
	if err != nil {
		return -1
	}
	d, err := strconv.Atoi(dL[2])
	if err != nil {
		return -1
	}
	sum := d
	for i := 0; i < m; i++ {
		sum += mDays[i]
	}
	if m > 2 && isLeepYear(y) {
		sum++
	}
	return sum
}

func isLeepYear(y int) bool {
	if (y%4 == 0 && y%100 != 0) || y%400 == 0 {
		return true
	}
	return false
}

func TestTest(t *testing.T) {
	for _, date := range []string{"2019-01-09", "2019-02-10", "2003-03-01", "2004-03-01"} {
		fmt.Println(dayOfYear(date))
	}
}
