package utils

import "time"

func Date() string {

	loc, _ := time.LoadLocation("Asia/Jakarta")

	date := time.Now().In(loc).Format("2006-01-02")
	return date
}

func DateTime() string {

	loc, _ := time.LoadLocation("Asia/Jakarta")

	date := time.Now().In(loc).Format("2006-01-02 15:04:05")
	return date
}
