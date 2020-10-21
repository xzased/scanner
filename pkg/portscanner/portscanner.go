package portscanner

import (
	"fmt"
	"net"
	"time"

	"github.com/rs/zerolog/log"
)

type ScanRequest struct {
	Host string
	Port int
}

type ScanResult struct {
	Host   string
	Port   int
	Open   bool
	Err    error
	Data string
}

func (sr *ScanRequest) HostPort() string {
	return fmt.Sprintf("%s:%d", sr.Host, sr.Port)
}

func (s *ScanResult) String() string {
	return fmt.Sprintf("%s:%d ok=%v err=%s data=%q",
		s.Host, s.Port, s.Open, s.Err, s.Data)
}

func ScanPort(sr ScanRequest) ScanResult {
	res := ScanResult{Host: sr.Host, Port: sr.Port}
	var data string
	hostport := sr.HostPort()
	log.Debug().Str("hostport", hostport).Msg("scanning")
	conn, err := net.DialTimeout("tcp", hostport, 2*time.Second)
	if err != nil {
		res.Err = err
		return res
	}
	defer conn.Close()
	dataBuffer := make([]byte, 1024)
	conn.SetDeadline(time.Now().Add(2*time.Second))
	n, err := conn.Read(dataBuffer)
	if err == nil {
		data = string(dataBuffer[:n])
	}
	res.Open = true
	res.Data = data
	return res
}