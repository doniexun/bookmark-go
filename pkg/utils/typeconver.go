package utils

import (
	"strconv"
)

func Int2str(i int) string {
	return strconv.Itoa(i)
}

func Str2int(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}
