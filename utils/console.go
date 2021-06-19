package utils

import (
	"fmt"
	"time"
)

func getTimeStr() string{
	return time.Now().Format("2006-01-02 15:04:05") + ": "
}

func ConsolePl(a ...interface{}){
	fmt.Println(getTimeStr(), a)
}

func ConsolePf(format string, a ...interface{}){
	format = getTimeStr() + format
	fmt.Printf(format, a...)
}
