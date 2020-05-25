package config

import (
	"os"
)

type ServerConfig struct {
	Host     string
	Port     int
	Timeouts map[string]uint16
	KeyFile  string
	CertFile string
}

func (sc *ServerConfig) CheckTlsFiles() error {
	_, err := os.OpenFile(sc.KeyFile, os.O_RDONLY, 0644)
	if os.IsNotExist(err) {
		return err
	}

	_, err = os.OpenFile(sc.CertFile, os.O_RDONLY, 0644)
	if os.IsNotExist(err) {
		return err
	}

	return nil
}
