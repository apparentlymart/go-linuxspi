// +build linux

// Package linuxspi provides access to the Linux "spidev" interface, providing
// access to SPI controllers from userspace.
package linuxspi

import (
	"fmt"
	"github.com/apparentlymart/go-spi/spi"
	"os"
	"syscall"
	"unsafe"
)

const (
	// Defining these here rather than referencing linux/spidev/spidev.h
	// because cgo breaks cross-compilation and cross-compilation is very
	// useful for embedded systems.
	// This will break if the ioctl numbers change in a future version of
	// Linux.
	ioctlWriteMode        uintptr = 0x40016b01
	ioctlWriteBitsPerWord uintptr = 0x40016b03
	ioctlWriteMaxSpeed    uintptr = 0x40046b04
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

func (dev spiDev) SetMaxSpeedHz(speed uint32) error {
	fmt.Println("Set max speed to", speed)
	file := dev.File
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, file.Fd(), ioctlWriteMaxSpeed, uintptr(unsafe.Pointer(&speed)))
	return err
}
