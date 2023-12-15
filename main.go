package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-daq/smbus"
)

func main() {
	var (
		err error
	)

	myHat := new(RGBHat)
	err = myHat.init()
	if err != nil {
		log.Fatalf("could not init device, err=[%s]", err)
	}

	for {
		err = myHat.updateFan()
		if err != nil {
			log.Printf("problem updating fan speed, err=[%s]", err)
		}

		time.Sleep(4 * time.Second)
	}

}

type RGBHat struct {
	i2CAddr      uint8
	bus          *smbus.Conn
	lastFanSpeed uint8
	nextFanSpeed uint8
	temp         float64
	powerOK      bool
}

// updateFan checks the current temp and updates the fan speed accordingly
func (h *RGBHat) updateFan() error {
	var err error = nil
	h.temp, err = getTemp()
	if err != nil {
		log.Printf("problem getting temp, err=[%s]", err)
	} else {
		log.Printf("temp is: %f, power ok: %t", h.temp, h.powerOK)
	}

	h.powerOK, err = powerOK()
	if err != nil {
		log.Printf("problem getting power status, err=[%s]", err)
	}

	if h.temp <= 36 {
		h.nextFanSpeed = FanOff
	} else if h.temp <= 40 {
		h.nextFanSpeed = Fan20
	} else if h.temp <= 45 {
		h.nextFanSpeed = Fan40
	} else if h.temp <= 50 {
		h.nextFanSpeed = Fan80
	} else if h.temp > 50 {
		h.nextFanSpeed = Fan100
	}

	// Check that we're not wasting time setting the fan to the speed it's already at
	if h.lastFanSpeed != h.nextFanSpeed {
		log.Printf("setting fan speed to: %d", h.nextFanSpeed)
		err = h.setFanSpeed()
	}
	return err
}

// init initializes the RGBHat struct
func (h *RGBHat) init() error {
	conn, err := smbus.Open(1, I2CAddr)
	if err != nil {
		return err
	}

	h.bus = conn
	h.i2CAddr = I2CAddr
	h.lastFanSpeed = FanInvalid
	h.powerOK, err = powerOK()
	if err != nil {
		log.Printf("non fatal error, getting initial power status, err=[%s]", err)
	}

	err = h.bus.WriteReg(h.i2CAddr, RegRGBMode, ModeBreathing)
	if err != nil {
		log.Printf("non fatal error, setting RGB scheme, err=[%s]", err)
	}

	err = h.bus.WriteReg(h.i2CAddr, RegRGBBreathScheme, BreathPurple)
	if err != nil {
		log.Printf("non fatal error, setting RGB colour, err=[%s]", err)
	}

	return nil
}

// setFanSpeed sets the fan speed register based on details on the RGBHat struct
func (h *RGBHat) setFanSpeed() error {
	err := h.bus.WriteReg(h.i2CAddr, RegFanSpeed, h.nextFanSpeed)
	if err != nil {
		return err
	}
	if !h.powerOK {
		log.Printf("alarm on pi power supply, disabling LEDs to reduce consumption")
		err = h.bus.WriteReg(h.i2CAddr, RegRGBOff, RGBOff)
		if err != nil {
			log.Printf("non fatal error, attempting to disable RGB LEDs, err=[%s]", err)
		}
	}
	h.lastFanSpeed = h.nextFanSpeed
	return nil
}

// getTemp gets the current CPU temp via /sys/class/thermal/thermal_zone0/temp
func getTemp() (float64, error) {
	tempInt, err := getSysInt(cpuTempHwMon)
	if err != nil {
		return 0, err
	}

	return float64(tempInt / 1000), nil
}

func powerOK() (bool, error) {
	alarm, err := getSysInt(powerHwMon)
	if err != nil {
		return false, err
	}

	if alarm == 0 {
		return true, nil
	}
	return false, nil
}

func getSysInt(path string) (int, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	fileContent = []byte(strings.TrimRight(string(fileContent), "\n"))

	sysInt, err := strconv.Atoi(string(fileContent))
	if err != nil {
		return 0, err
	}
	return sysInt, nil
}
