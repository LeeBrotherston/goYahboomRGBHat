# Yahboom RGB Hat

## Background

The Yahboom RGB Hat for Raspberry Pi does not come with any software, and the examples provided are very basic, sometimes referring to outdated libraries, etc.  Additionally the examples (and subsequently a number of opensource tools written for this hat) execute a shell command to obtain the temperature, and then a the smbus library for python which will regularly panic the kernel on a Pi due to an implementation error on the I2C bus.

So I wrote this, it's not at all special, but it runs without crashing my Pi and does not rely on external binaries.

## Building

Assuming you have `go` installed you should just be able to:

`go build`

Currently there is one dependancy, [github.com/go-daq/smbus](github.com/go-daq/smbus) which is used to access the smbus/i2c bus.  This should be automatically downloaded as part of the `go build`

Then you can simply run the binary produced.

The RGBLEDs are set to a purple breathing scheme to indicate that it has started working

## Stuff I didn't do

- I didn't write an install script, but you can copy/run the binary wherever, it's self contained and doesn't take a config currently
- No config file... Parsing the config would be longer than the whole of the rest, and I only really wrote this for me, but I designed it to be super simple to make changes in line with what you want
- I left all the contants needed to support the RGB LEDs, but haven't really done anything much with them at this point.
