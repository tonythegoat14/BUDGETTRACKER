package main

import (
	"time"
	"log"
)

var Time = func() time.Time {
	return time.Now()
}


func RestoreTime(oldTime func() time.Time) {
	log.Println("Restoring time")
	Time = oldTime
}

func MockTime(now ring) {
	log.Println("Mocking time to", now)
	Time = func() time.Time {
		mockTime, err := time.Parse("2006-01-02 15:04:05", now)
		if err != nil {
			log.Fatal("Unable to parse time", err)
		}
		return mockTime
	}
}
