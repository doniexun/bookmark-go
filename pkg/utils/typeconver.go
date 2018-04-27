package utils

import (
	"strconv"
)

func Int2str(i int) string {
	return strconv.Itoa(i)
}

// string无法转换成int时返回 defaultval
func Str2int(s string, defaultval int) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		return defaultval
	}
	return num
}
