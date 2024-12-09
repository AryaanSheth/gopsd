package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/AryaanSheth/gopsd"
)

func main() {
	var gps *gopsd.Session
	var err error

	address := "localhost:2947"

	if gps, err = gopsd.Dial(address); err != nil {
		panic(fmt.Sprintf("Failed to connect to GPSD: %s", err))
	}
	defer gps.Close()

	gps.AddFilter("TPV", func(data []byte) {
		var tpv gopsd.TPV
		if err := json.Unmarshal(data, &tpv); err == nil {
			fmt.Printf("TPV - Mode: %d, Time: %v\n", tpv.Mode, tpv.Time)
		}
	})

	gps.AddFilter("SKY", func(data []byte) {
		var sky gopsd.SKY
		if err := json.Unmarshal(data, &sky); err == nil {
			fmt.Printf("SKY - %d satellites\n", len(sky.Satellites))
		}
	})

	done := gps.Watch()

	select {
	case <-done:
		fmt.Println("GPSD connection closed")
	case <-time.After(5 * time.Minute):
		fmt.Println("Connection timeout")
		gps.Close()
	}
}
