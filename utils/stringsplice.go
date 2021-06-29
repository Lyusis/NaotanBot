package utils

import "strconv"

func SingleFront(msg0 string, num int) string {
	return strconv.Itoa(num) + msg0
}

func SingleBack(msg0 string, num int) string {
	return msg0 + strconv.Itoa(num)
}

func Double(msg0, msg1 string, num int) string {
	return msg0 + strconv.Itoa(num) + msg1
}
