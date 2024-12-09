/*
* interface.go
*
* Interfaces derived from https://gpsd.gitlab.io/gpsd/gpsd_json.html
*
 */

package gopsd

import (
	"bufio"
	"net"
	"sync"
)

// Constants
const (
	// mode constants
	NoValue Mode = 0
	NoFix   Mode = 1
	Mode2D  Mode = 2
	Mode3D  Mode = 3

	DefaultAddress = "localhost:2947"
)

// Interfaces

type Filter func(interface{}) // GPSD Server Filter function

type Mode byte // Fix Mode (0: No Value, 1: No Fix, 2: 2D, 3: 3D)

type Session struct {
	conn    net.Conn      // Client GPSD Server connection
	reader  *bufio.Reader // Client GPSD Server reader
	filters sync.Map      // Filters for the GPSD Server
}

type gopsdReport struct {
	Class string `json:"class"` // Type of report
}

type TPV struct {
	Class       string  `json:"class"`                 // Fixed: "TPV"
	Device      string  `json:"device,omitempty"`      // Name of the originating device.
	Mode        int     `json:"mode"`                  // The mode of operation
	Alt         float64 `json:"alt,omitempty"`         // Altitude (meters)
	AltHAE      float64 `json:"altHAE,omitempty"`      // Altitude above ellipsoid (meters)
	AltMSL      float64 `json:"altMSL,omitempty"`      // Altitude above mean sea level (meters)
	Ant         int     `json:"ant,omitempty"`         // Antenna status
	Climb       float64 `json:"climb,omitempty"`       // Climb rate (meters per second)
	ClockBias   float64 `json:"clockbias,omitempty"`   // Clock bias (seconds)
	ClockDrift  float64 `json:"clockdrift,omitempty"`  // Clock drift (seconds per second)
	Datum       string  `json:"datum,omitempty"`       // Datum used for coordinates
	Depth       float64 `json:"depth,omitempty"`       // Depth (meters)
	DgpsAge     float64 `json:"dgpsAge,omitempty"`     // DGPS age (seconds)
	DgpsSta     float64 `json:"dgpsSta,omitempty"`     // DGPS station ID
	Ecefx       float64 `json:"ecefx,omitempty"`       // ECEF X coordinate (meters)
	Ecefy       float64 `json:"ecefy,omitempty"`       // ECEF Y coordinate (meters)
	Ecefz       float64 `json:"ecefz,omitempty"`       // ECEF Z coordinate (meters)
	EcefpAcc    float64 `json:"ecefpAcc,omitempty"`    // ECEF position accuracy (meters)
	Ecefvx      float64 `json:"ecefvx,omitempty"`      // ECEF velocity X component (meters per second)
	Ecefvy      float64 `json:"ecefvy,omitempty"`      // ECEF velocity Y component (meters per second)
	Ecefvz      float64 `json:"ecefvz,omitempty"`      // ECEF velocity Z component (meters per second)
	EcefvAcc    float64 `json:"ecefvAcc,omitempty"`    // ECEF velocity accuracy (meters per second)
	Epc         float64 `json:"epc,omitempty"`         // Ephemeris clock bias (seconds)
	Epd         float64 `json:"epd,omitempty"`         // Ephemeris clock drift (seconds per second)
	Eph         float64 `json:"eph,omitempty"`         // Ephemeris health status
	Eps         float64 `json:"eps,omitempty"`         // Ephemeris signal strength (dB)
	Ept         float64 `json:"ept"`                   // Ephemeris time (seconds)
	Epx         float64 `json:"epx,omitempty"`         // Ephemeris X position (meters)
	Epy         float64 `json:"epy,omitempty"`         // Ephemeris Y position (meters)
	Epv         float64 `json:"epv,omitempty"`         // Ephemeris Z position (meters)
	GeoidSep    float64 `json:"geoidSep,omitempty"`    // Geoid separation (meters)
	Jam         int     `json:"jam,omitempty"`         // Jam status
	Lat         float64 `json:"lat,omitempty"`         // Latitude (degrees)
	LeapSeconds int     `json:"leapseconds,omitempty"` // Leap seconds correction
	Lon         float64 `json:"lon,omitempty"`         // Longitude (degrees)
	MagTrack    float64 `json:"magtrack,omitempty"`    // Magnetic track (degrees)
	MagVar      float64 `json:"magvar,omitempty"`      // Magnetic variation (degrees)
	RelD        float64 `json:"relD,omitempty"`        // Relative distance (meters)
	RelE        float64 `json:"relE,omitempty"`        // Relative east (meters)
	RelN        float64 `json:"relN,omitempty"`        // Relative north (meters)
	Sep         float64 `json:"sep,omitempty"`         // Separation distance (meters)
	Speed       float64 `json:"speed,omitempty"`       // Speed (meters per second)
	Status      int     `json:"status,omitempty"`      // Status of the GPS fix
	Temp        float64 `json:"temp,omitempty"`        // Temperature (degrees Celsius)
	Time        string  `json:"time,omitempty"`        // Timestamp of the fix
	Track       float64 `json:"track,omitempty"`       // Heading (degrees)
	VelD        float64 `json:"velD,omitempty"`        // Vertical velocity (meters per second)
	VelE        float64 `json:"velE,omitempty"`        // East velocity (meters per second)
	VelN        float64 `json:"velN,omitempty"`        // North velocity (meters per second)
	Wanglem     float64 `json:"wanglem,omitempty"`     // Longitude of the antenna (degrees)
	Wangler     float64 `json:"wangler,omitempty"`     // Latitude of the antenna (degrees)
	Wanglet     float64 `json:"wanglet,omitempty"`     // Altitude of the antenna (meters)
	Wspeedr     float64 `json:"wspeedr,omitempty"`     // Relative wind speed (meters per second)
}

type Satellite struct {
	PRN    int     `json:"PRN"`              // PRN ID of the satellite
	Az     float64 `json:"az"`               // Azimuth, degrees from true north
	El     float64 `json:"el"`               // Elevation in degrees
	FreqID int     `json:"freqid,omitempty"` // For GLONASS: the frequency ID of the signal
	GNSSID int     `json:"gnssid,omitempty"` // The GNSS ID
	Health int     `json:"health,omitempty"` // Health of the satellite (0=unknown, 1=OK, 2=unhealthy)
	SS     float64 `json:"ss,omitempty"`     // Signal to Noise ratio in dBHz
	SigID  int     `json:"sigid,omitempty"`  // Signal ID of this signal
	SVID   int     `json:"svid,omitempty"`   // Satellite ID within its constellation
	Used   bool    `json:"used"`             // Used in current solution
}

type SKY struct {
	Class      string      `json:"class"`                // Fixed: "SKY"
	Device     string      `json:"device,omitempty"`     // Name of originating device
	NSat       int         `json:"nSat,omitempty"`       // Number of satellites in the sky
	GDop       float64     `json:"gdop,omitempty"`       // Geometric dilution of precision
	HDop       float64     `json:"hdop,omitempty"`       // Horizontal dilution of precision
	PDop       float64     `json:"pdop,omitempty"`       // Position dilution of precision
	Pr         float64     `json:"pr,omitempty"`         // Pseudorange in meters
	PrRate     float64     `json:"prRate,omitempty"`     // Pseudorange Rate of Change
	PrRes      float64     `json:"prRes,omitempty"`      // Pseudorange residue in meters
	Qual       int         `json:"qual,omitempty"`       // Quality Indicator
	Satellites []Satellite `json:"satellites,omitempty"` // List of satellite objects
	Tdop       float64     `json:"tdop,omitempty"`       // Time dilution of precision
	Time       string      `json:"time,omitempty"`       // Time/date stamp in ISO8601 format
	USat       int         `json:"uSat,omitempty"`       // Number of satellites used in navigation
	VDop       float64     `json:"vdop,omitempty"`       // Vertical dilution of precision
	XDop       float64     `json:"xdop,omitempty"`       // Longitudinal dilution of precision
	YDop       float64     `json:"ydop,omitempty"`       // Latitudinal dilution of precision
}

type GST struct {
	Class  string  `json:"class"`            // Fixed: "GST"
	Device string  `json:"device,omitempty"` // Name of originating device
	Time   string  `json:"time,omitempty"`   // Time/date stamp in ISO8601 format, UTC
	RMS    float64 `json:"rms,omitempty"`    // Standard deviation of range inputs to the navigation process
	Major  float64 `json:"major,omitempty"`  // Standard deviation of semi-major axis of error ellipse (meters)
	Minor  float64 `json:"minor,omitempty"`  // Standard deviation of semi-minor axis of error ellipse (meters)
	Orient float64 `json:"orient,omitempty"` // Orientation of semi-major axis of error ellipse (degrees from true north)
	Alt    float64 `json:"alt,omitempty"`    // Standard deviation of altitude error (meters)
	Lat    float64 `json:"lat,omitempty"`    // Standard deviation of latitude error (meters)
	Lon    float64 `json:"lon,omitempty"`    // Standard deviation of longitude error (meters)
	VE     float64 `json:"ve,omitempty"`     // Standard deviation of East velocity error (meters/second)
	VN     float64 `json:"vn,omitempty"`     // Standard deviation of North velocity error (meters/second)
	VU     float64 `json:"vu,omitempty"`     // Standard deviation of Up velocity error (meters/second)
}

type ATT struct {
	Class    string  `json:"class"`              // Fixed: "ATT"
	Device   string  `json:"device"`             // Name of originating device
	Time     string  `json:"time,omitempty"`     // Time/date stamp in ISO8601 format, UTC
	TimeTag  string  `json:"timeTag,omitempty"`  // Arbitrary time tag of measurement
	Heading  float64 `json:"heading,omitempty"`  // Heading, degrees from true north
	MagSt    string  `json:"mag_st,omitempty"`   // Magnetometer status
	MHeading float64 `json:"mheading,omitempty"` // Heading, degrees from magnetic north
	Pitch    float64 `json:"pitch,omitempty"`    // Pitch in degrees
	PitchSt  string  `json:"pitch_st,omitempty"` // Pitch sensor status
	Rot      float64 `json:"rot,omitempty"`      // Rate of Turn in degrees per minute
	Yaw      float64 `json:"yaw,omitempty"`      // Yaw in degrees
	YawSt    string  `json:"yaw_st,omitempty"`   // Yaw sensor status
	Roll     float64 `json:"roll,omitempty"`     // Roll in degrees
	RollSt   string  `json:"roll_st,omitempty"`  // Roll sensor status
	Dip      float64 `json:"dip,omitempty"`      // Local magnetic inclination, degrees
	MagLen   float64 `json:"mag_len,omitempty"`  // Scalar magnetic field strength
	MagX     float64 `json:"mag_x,omitempty"`    // X component of magnetic field strength
	MagY     float64 `json:"mag_y,omitempty"`    // Y component of magnetic field strength
	MagZ     float64 `json:"mag_z,omitempty"`    // Z component of magnetic field strength
	AccLen   float64 `json:"acc_len,omitempty"`  // Scalar acceleration
	AccX     float64 `json:"acc_x,omitempty"`    // X component of acceleration (m/s^2)
	AccY     float64 `json:"acc_y,omitempty"`    // Y component of acceleration (m/s^2)
	AccZ     float64 `json:"acc_z,omitempty"`    // Z component of acceleration (m/s^2)
	GyroX    float64 `json:"gyro_x,omitempty"`   // X component of angular rate (deg/s)
	GyroY    float64 `json:"gyro_y,omitempty"`   // Y component of angular rate (deg/s)
	GyroZ    float64 `json:"gyro_z,omitempty"`   // Z component of angular rate (deg/s)
	Depth    float64 `json:"depth,omitempty"`    // Water depth in meters
	Temp     float64 `json:"temp,omitempty"`     // Temperature at the sensor (°C)
}

type TOFF struct {
	Class     string  `json:"class"`      // Fixed: "TOFF"
	Device    string  `json:"device"`     // Name of the originating device
	RealSec   float64 `json:"real_sec"`   // Seconds from the GPS clock
	RealNSec  float64 `json:"real_nsec"`  // Nanoseconds from the GPS clock
	ClockSec  float64 `json:"clock_sec"`  // Seconds from the system clock
	ClockNSec float64 `json:"clock_nsec"` // Nanoseconds from the system clock
}

type PPS struct {
	Class     string   `json:"class"`          // Fixed: "PPS"
	Device    string   `json:"device"`         // Name of the originating device
	RealSec   float64  `json:"real_sec"`       // Seconds from the PPS source
	RealNSec  float64  `json:"real_nsec"`      // Nanoseconds from the PPS source
	ClockSec  float64  `json:"clock_sec"`      // Seconds from the system clock
	ClockNSec float64  `json:"clock_nsec"`     // Nanoseconds from the system clock
	Precision float64  `json:"precision"`      // NTP-style estimate of PPS precision
	Shm       string   `json:"shm"`            // shm key of this PPS
	QErr      *float64 `json:"qErr,omitempty"` // Quantization error of the PPS, in picoseconds (optional)
}

type OSC struct {
	Class       string  `json:"class"`       // Fixed: "OSC"
	Device      string  `json:"device"`      // Name of the originating device
	Running     bool    `json:"running"`     // Indicates if the oscillator is currently running
	Reference   bool    `json:"reference"`   // Indicates if the oscillator is receiving a GPS PPS signal
	Disciplined bool    `json:"disciplined"` // Indicates if the GPS PPS signal is sufficiently stable and disciplining the local oscillator
	Delta       float64 `json:"delta"`       // Time difference (in nanoseconds) between PPS output and the most recent GPS PPS input
}

type DEVICES struct {
	Class   string   `json:"class"`            // Fixed: "DEVICES"
	Devices []DEVICE `json:"devices"`          // List of device descriptions
	Remote  *string  `json:"remote,omitempty"` // URL of the remote daemon (optional)
}

type DEVICE struct {
	Class     string  `json:"class"`               // Fixed: "DEVICE"
	Activated string  `json:"activated,omitempty"` // Time the device was activated, if inactive this field is absent
	Bps       int     `json:"bps,omitempty"`       // Device speed in bits per second
	Cycle     float64 `json:"cycle,omitempty"`     // Device cycle time in seconds
	Driver    string  `json:"driver,omitempty"`    // GPSD’s name for the device driver type
	Flags     int     `json:"flags,omitempty"`     // Bit vector of property flags
	Hexdata   string  `json:"hexdata,omitempty"`   // Data to send to the GNSS receiver in hexadecimal
	Mincycle  float64 `json:"mincycle,omitempty"`  // Minimum cycle time in seconds (read-only)
	Native    int     `json:"native,omitempty"`    // NMEA mode (0) or alternate mode (1)
	Parity    string  `json:"parity,omitempty"`    // Parity: N, O or E (No parity, Odd, Even)
	Path      string  `json:"path,omitempty"`      // Device path
	Readonly  bool    `json:"readonly,omitempty"`  // True if device is read-only
	Sernum    string  `json:"sernum,omitempty"`    // Hardware serial number
	Stopbits  int     `json:"stopbits"`            // Stop bits (1 or 2)
	Subtype   string  `json:"subtype,omitempty"`   // Version information
	Subtype1  string  `json:"subtype1,omitempty"`  // Additional version information
}

type WATCH struct {
	Class   string  `json:"class"`             // Fixed: "WATCH"
	Enable  *bool   `json:"enable,omitempty"`  // Enable/disable watcher mode (default true)
	JSON    *bool   `json:"json,omitempty"`    // Enable/disable JSON reports (default false)
	NMEA    *bool   `json:"nmea,omitempty"`    // Enable/disable pseudo-NMEA dumping (default false)
	Raw     *int    `json:"raw,omitempty"`     // Controls raw mode (1: hex-dump, 2: verbatim)
	Scaled  *bool   `json:"scaled,omitempty"`  // Apply scaling divisors (default false)
	Split24 *bool   `json:"split24,omitempty"` // Aggregate AIS type24 parts (default false)
	PPS     *bool   `json:"pps,omitempty"`     // Emit TOFF/PPS messages (default false)
	Device  *string `json:"device,omitempty"`  // Watch only the specified device
	Remote  *string `json:"remote,omitempty"`  // URL of the remote daemon (optional)
	Timing  *bool   `json:"timing,omitempty"`  // Undocumented; developer use only
}

type POLL struct {
	Class  string `json:"class"`  // Fixed: "POLL"
	Time   string `json:"time"`   // Timestamp in ISO 8601 format
	Active int    `json:"active"` // Count of active devices
	TPV    []TPV  `json:"tpv"`    // List of TPV objects
	Sky    []SKY  `json:"sky"`    // List of SKY objects
}

type ERROR struct {
	Class   string `json:"class"`   // Fixed: "ERROR"
	Message string `json:"message"` // Textual error message
}
