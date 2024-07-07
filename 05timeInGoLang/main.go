package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time Management in go lang")
	// we have a time module to deal with time in go Lang
	// to get the current time
	timeNow := time.Now()
	fmt.Println(`current time is: `, timeNow)

	// formatted time
	/*
		You can use any variation below in anyorder
		You can also remove any field (date , day ,am/pm,time,min, sec etc) based on your need
	*/
	// timeFormatted := time.Now().Format("01-02-2006 Monday")
	// timeFormatted := time.Now().Format("01-02-2006 Monday 15:04:05")
	// timeFormatted := time.Now().Format("01-02-2006 15:04:05 Mon")
	// timeFormatted := time.Now().Format("01/02/2006 15:04:05 Monday")
	// timeFormatted := time.Now().Format("01 Jan 2006 15:04:05 Monday")
	// timeFormatted := time.Now().Format("01/02/2006 15:04 pm Monday")
	timeFormatted := time.Now().Format("01/02/2006 15:04:05 pm Monday")
	// timeFormatted := time.Now().Format("01 Jan 2006 15:04:05 Monday")

	fmt.Println("Formatted time: ", timeFormatted)

	// creating our own date
	createdDate := time.Date(2020, time.February, 28, 15, 20, 45, 0, time.UTC) // all these arguments needs to be passed
	//args are: year, month, date, hour, min, sec, nanosec,timeZone location

	fmt.Println("Created date is : ", createdDate.Format("01-02-2006 Mon 15:04:05"))

}
