package main

import (
	"fmt"
	"time"

	"github.com/AryaanSheth/gopsd"
)

func main() {
	var gps *gopsd.Session
	var err error

	address := "localhost:2947"

	fmt.Printf("Connecting to GPSD at %s\n", address)
	if gps, err = gopsd.Dial(address); err != nil {
		panic(fmt.Sprintf("Failed to connect to GPSD: %s", err))
	}
	defer gps.Close()
	fmt.Println("Connected to GPSD")

	fmt.Println("Adding filters")
	gps.AddFilter("TPV", func(data interface{}) {
		if tpv, ok := data.(*gopsd.TPV); ok {
			fmt.Printf("TPV - Mode: %d, Time: %v\n", tpv.Mode, tpv.Time)
		} else {
			fmt.Printf("TPV - %v\n", data)
		}
	})

	gps.AddFilter("SKY", func(data interface{}) {
		if sky, ok := data.(*gopsd.SKY); ok {
			fmt.Printf("SKY - %d satellites\n", len(sky.Satellites))
		} else {
			fmt.Printf("SKY - %v\n", data)
		}
	})
	fmt.Println("Filters added")

	fmt.Println("Starting watch")
	done := gps.Watch()
	fmt.Println("Watch started")

	select {
	case <-done:
		fmt.Println("GPSD connection closed")
	case <-time.After(30 * time.Second):
		fmt.Println("Connection timeout")
		gps.Close()
	}
}
