package main

// Just a bunch of constants to make the code more readable when setting
// register values, values from:
// http://www.yahboom.net/xiazai/Raspberry_Pi_RGB_Cooling_HAT/I2C%20Communication%20Protocol.pdf
const I2CAddr = 0x0d

// RegisterNumbers
const (
	// RGB Values
	RegSelectControl = 0x00
	RegRedValue      = 0x01
	RegGreenValue    = 0x02
	RegBlueValue     = 0x03

	// RGB Effects
	RegRGBMode         = 0x04
	RegRGBSpeed        = 0x05
	RegRGBBreathScheme = 0x06
	RegRGBOff          = 0x07

	// Fan
	RegFanSpeed = 0x08
)

// Select Control
const (
	ControlRGB1   = 0x00
	ControlRGB2   = 0x01
	ControlRGB3   = 0x02
	ControlRGBAll = 0xFF
)

// RGB Mode
const (
	ModeWater     = 0x00
	ModeBreathing = 0x01
	ModeMarquee   = 0x02
	ModeRainbow   = 0x03
	ModeColourful = 0x04
)

// RGB Speed
const (
	RGBSpeedSlow = 0x01
	RGBSpeedMid  = 0x02
	RGBSpeedFast = 0x03
)

// RGB Breathing Colour
const (
	BreathRed    = 0x00
	BreathGreen  = 0x01
	BreathBlue   = 0x02
	BreathYellow = 0x03
	BreathPurple = 0x04
	BreathCyan   = 0x05
	BreathWhite  = 0x06
)

// Turn off
const (
	RGBOff = 0x01
)

const (
	FanOff     = 0x00
	Fan100     = 0x01
	Fan20      = 0x02
	Fan30      = 0x03
	Fan40      = 0x04
	Fan50      = 0x05
	Fan60      = 0x06
	Fan70      = 0x07
	Fan80      = 0x08
	Fan90      = 0x09
	FanInvalid = 0x10
)
