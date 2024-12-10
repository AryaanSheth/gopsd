package examples

import (
	"fmt"
	"log"
	"time"

	"github.com/AryaanSheth/gopsd"
)

type CustomGPSData struct {
	Class  string  `json:"class"`
	Device string  `json:"device"`
	Status int     `json:"status,omitempty"`
	Ept    float64 `json:"ept,omitempty"`
	Lat    float64 `json:"lat,omitempty"`
	Lon    float64 `json:"lon,omitempty"`
	Alt    float64 `json:"alt,omitempty"`
	Rms    float64 `json:"rms,omitempty"`
	Major  float64 `json:"major,omitempty"`
	Minor  float64 `json:"minor,omitempty"`
	Orient float64 `json:"orient,omitempty"`
}

// Example: Populate a Custom Struct
func customStruct() {
	gps, err := gopsd.Dial(gopsd.DefaultAddress)
	if err != nil {
		log.Fatalf("Failed to connect to GPSD: %v", err)
	}
	defer gps.Close()

	var tpv *gopsd.TPV
	var gst *gopsd.GST

	gps.AddFilter("TPV", func(r interface{}) { tpv = r.(*gopsd.TPV) })
	gps.AddFilter("GST", func(r interface{}) { gst = r.(*gopsd.GST) })

	done := gps.Watch()
	go func() {
		<-done
	}()
	time.Sleep(3 * time.Second)

	customData := CustomGPSData{}

	if tpv != nil {
		customData.Class = tpv.Class
		customData.Device = tpv.Device
		customData.Status = tpv.Status
		customData.Ept = tpv.Ept
		customData.Lat = tpv.Lat
		customData.Lon = tpv.Lon
		customData.Alt = tpv.Alt
	}

	if gst != nil {
		customData.Rms = gst.RMS
		customData.Major = gst.Major
		customData.Minor = gst.Minor
		customData.Orient = gst.Orient
	}

	fmt.Printf("Custom GPS Data: %+v\n", customData)
}
