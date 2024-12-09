package examples

import (
	"fmt"
	"time"

	"github.com/AryaanSheth/gopsd"
)

// Example: Adding Multiple Filters
func multipleFilters() {
	var gps *gopsd.Session
	var err error

	if gps, err = gopsd.Dial(gopsd.DefaultAddress); err != nil {
		panic(fmt.Sprintf("Failed to connect to GPSD: %s", err))
	}
	defer gps.Close()

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

	done := gps.Watch()

	select {
	case <-done:
		fmt.Println("GPSD connection closed")
	case <-time.After(30 * time.Second):
		fmt.Println("Connection timeout")
		gps.Close()
	}
}
