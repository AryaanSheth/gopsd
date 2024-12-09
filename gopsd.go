package gopsd

import (
	"bufio"
	"errors"
	"net"
	"sync"
	"time"

	"github.com/bytedance/sonic"
)

const (
	syscallBufferSize = 4096
	connectionType    = "tcp4"
)

// Open a new connection to the GPSD daemon
func Dial(address string) (*Session, error) {
	return dialCommon(net.Dial(connectionType, address))
}

// Open a new connection to GPSD with timeout
func DialTimeout(address string, to time.Duration) (*Session, error) {
	return dialCommon(net.DialTimeout(connectionType, address, to))
}

// Create a new session with a connection and reader
func dialCommon(c net.Conn, err error) (*Session, error) {
	if err != nil {
		return nil, err
	}

	session := &Session{
		conn:    c,
		reader:  bufio.NewReaderSize(c, syscallBufferSize),
		filters: sync.Map{},
	}

	// Read initial connection message without alloc
	_, _ = session.reader.ReadString('\n')

	return session, nil
}

// GPSD watcher session for checking reports
func (s *Session) Watch() <-chan bool {
	_, _ = s.conn.Write([]byte(`?WATCH={"enable":true,"json":true}`))

	done := make(chan bool, 1)
	go s.watchReports(done)
	return done
}

// Send a command to GPSD
func (s *Session) SendCommand(command string) {
	_, _ = s.conn.Write([]byte("?" + command + ";"))
}

// Attach a filter to a class of reports
func (s *Session) AddFilter(class string, f Filter) {
	filters, _ := s.filters.LoadOrStore(class, []Filter{})
	s.filters.Store(class, append(filters.([]Filter), f))
}

// Safely close the GPSD connection
func (s *Session) Close() error {
	if s.conn == nil {
		return errors.New("GPSD socket is already closed")
	}
	return s.conn.Close()
}

// Handle report watching and dispatching
func (s *Session) watchReports(done chan<- bool) {
	defer func() { done <- true }()

	scanner := bufio.NewScanner(s.reader)
	scanner.Buffer(make([]byte, syscallBufferSize), syscallBufferSize*10)

	for scanner.Scan() {
		lineBytes := scanner.Bytes()

		var reportPeek gopsdReport
		if err := sonic.Unmarshal(lineBytes, &reportPeek); err != nil {
			continue
		}

		if report := s.unmarshalReport(reportPeek.Class, lineBytes); report != nil {
			if filtersRaw, ok := s.filters.Load(reportPeek.Class); ok {
				s.dispatchReport(reportPeek.Class, report, filtersRaw.([]Filter))
			}
		}
	}
}

// Convert a report to a struct
func (s *Session) unmarshalReport(class string, data []byte) interface{} {
	var report interface{}
	switch class {
	case "TPV":
		var r TPV
		if err := sonic.Unmarshal(data, &r); err == nil {
			report = &r
		}
	case "SKY":
		var r SKY
		if err := sonic.Unmarshal(data, &r); err == nil {
			report = &r
		}
	case "GST":
		var r GST
		if err := sonic.Unmarshal(data, &r); err == nil {
			report = &r
		}
	case "ATT":
		var r ATT
		if err := sonic.Unmarshal(data, &r); err == nil {
			report = &r
		}
	case "DEVICES":
		var r DEVICES
		if err := sonic.Unmarshal(data, &r); err == nil {
			report = &r
		}
	case "PPS":
		var r PPS
		if err := sonic.Unmarshal(data, &r); err == nil {
			report = &r
		}
	case "TOFF":
		var r TOFF
		if err := sonic.Unmarshal(data, &r); err == nil {
			report = &r
		}
	case "ERROR":
		var r ERROR
		if err := sonic.Unmarshal(data, &r); err == nil {
			report = &r
		}
	}
	return report
}

// Call all filters for a class
func (s *Session) dispatchReport(class string, report interface{}, filters []Filter) {
	for _, f := range filters {
		f(report)
	}
}
