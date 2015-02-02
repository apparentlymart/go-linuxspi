// +build linux

// Package linuxspi provides access to the Linux "spidev" interface, providing
// access to SPI controllers from userspace.
package linuxspi

import (
	"fmt"
	"github.com/apparentlymart/go-spi/spi"
	"os"
)

type spiDev struct {
	// We embed *os.File here to automatically inherit its io.Reader and
	// io.Writer implementations, which are needed for the spi.Device
	// interface.
	*os.File
}

func OpenDevice(busNumber uint, chipSelectNumber uint) (spi.Device, error) {
	filename := fmt.Sprintf("/dev/spidev%d.%d", busNumber, chipSelectNumber)

	var file *os.File
	file, err := os.OpenFile(filename, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	return &spiDev{file}, nil
}

func (dev spiDev) Exchange(outData []byte, inData []byte) (int, error) {
	return 0, nil
}

func (dev spiDev) Request(outData []byte, inData []byte) (int, error) {
	return 0, nil
}

func (dev spiDev) SetBitOrder(order spi.BitOrder) error {
	return nil
}

func (dev spiDev) SetMode(mode spi.Mode) error {
	return nil
}
