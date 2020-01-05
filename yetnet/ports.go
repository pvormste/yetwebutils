package yetnet

import (
	"fmt"
	"net"
)

func IsPortOpen(port int) (bool, error) {
	inUse, err := IsPortInUse(port)
	return !inUse, err
}

func IsPortInUse(port int) (bool, error) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return true, nil
	}

	if err := l.Close(); err != nil {
		return false, err
	}

	return false, nil
}
